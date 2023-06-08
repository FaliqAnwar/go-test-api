package logzap

import (
	"log"
	"os"
	"strings"
	"time"

	"go-test-api/internal/model"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

/*
LogLevel:
debug
info
warn
error
dpanic
panic
fatal
*/

/*
logOption:
file
fileconsole
console
*/

const (
	defaultMaxSize = 1024
	sep            = string(os.PathSeparator)
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	env         model.Environment
)

func InitLogger(conf model.Config) (*zap.Logger, *zap.SugaredLogger) {
	env = model.StringToEnvironment(conf.App.Env)
	encoderConfig := getEncoder(model.EnvironmentToString(env))

	// set log level
	atomicLevel, err := zap.ParseAtomicLevel(conf.App.LogLevel)
	if err != nil {
		log.Printf("error during set log level: %s, will use default value", err)
		atomicLevel = zap.NewAtomicLevel()
	}

	hook := &lumberjack.Logger{
		Filename:   getFileName(conf.App.Name), // Log name
		MaxSize:    defaultMaxSize,             // File content size, MB
		MaxBackups: 7,                          // Maximum number of old files retained
		MaxAge:     100,                        // Maximum number of days to keep old files
		Compress:   true,                       // Is the file compressed
	}
	var writes = []zapcore.WriteSyncer{}

	switch strings.ToLower(conf.App.LogOption) {
	case "file":
		writes = append(writes, zapcore.AddSync(hook))

	case "fileconsole":
		writes = append(writes, zapcore.AddSync(os.Stdout))
		writes = append(writes, zapcore.AddSync(hook))

	case "console":
		writes = append(writes, zapcore.AddSync(os.Stdout))

	default:
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	opts := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Fields(zap.String("applicationName", conf.App.Name)),
	}

	// print stack trace for every error in dev environment
	if env == model.DEV_ENV || env == model.LOCAL_ENV {
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel), zap.Development())
	}

	// construction log
	logger = zap.New(core, opts...)
	sugarLogger = logger.Sugar()

	logger.Info("log initialized successfully")

	return logger, sugarLogger
}

func InitLoggerTest() {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	logger = observedLogger
	sugarLogger = logger.Sugar()
}

func getEncoder(env string) zapcore.EncoderConfig {
	var encoderConfig zapcore.EncoderConfig

	switch strings.ToLower(env) {
	case "prod":
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig = zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "severity",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLevel(),
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig = zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "severity",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLevel(),
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}
	}
	return encoderConfig
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

func getFileName(appName string) string {
	hostName, _ := os.Hostname() // nolint
	if hostName == "" {
		hostName = "log.log"
	}

	suffix := time.Now().Format("2006-01-02T15-04-05.000000000")
	fileName := "storages" + sep + "logs" + sep + "log_" + suffix + "_" + appName + ".log"

	return fileName
}
