package logzap

import (
	"context"
	"fmt"

	"go-test-api/internal/ctxdata"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	traceKey        = "logging.googleapis.com/trace"
	spanKey         = "logging.googleapis.com/spanId"
	traceSampledKey = "logging.googleapis.com/trace_sampled"
)

func setFieldFromContext(ctx context.Context) []zap.Field {
	if ctx != nil {
		fields := []zapcore.Field{
			zap.String("traceparent", ctxdata.GetTraceParent(ctx)),
			zap.String(traceKey, ctxdata.GetTraceID(ctx)),
			zap.String(spanKey, ctxdata.GetSpanID(ctx)),
			zap.Bool(traceSampledKey, ctxdata.GetTraceSampled(ctx)),
			zap.String("referenceNumber", ctxdata.GetReferenceNumber(ctx)),
			zap.String("host", ctxdata.GetHost(ctx)),
			zap.String("real_ip", ctxdata.GetRealIP(ctx)),
			zap.String("user_agent", ctxdata.GetUserAgent(ctx)),
			zap.String("pid", ctxdata.GetPid(ctx)),
			zap.String("path", ctxdata.GetPath(ctx)),
			zap.String("method", ctxdata.GetMethod(ctx)),
			zap.Int("status", ctxdata.GetStatus(ctx)),
		}
		return fields
	}
	return []zap.Field{}
}

func Info(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Info(message, fields...)
}

func Infof(ctx context.Context, message string, i ...interface{}) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Info(fmt.Sprintf(message, i...))
}

func Debug(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Debug(message, fields...)
}

func Debugf(ctx context.Context, message string, i ...interface{}) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Debug(fmt.Sprintf(message, i...))
}

func Error(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Error(message, fields...)
}

func Errorf(ctx context.Context, message string, i ...interface{}) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Error(fmt.Sprintf(message, i...))
}

func Warn(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Warn(message, fields...)
}

func Warnf(ctx context.Context, message string, i ...interface{}) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Warn(fmt.Sprintf(message, i...))
}

func Panic(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Panic(message, fields...)
}

func DPanic(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).DPanic(message, fields...)
}

func Fatal(ctx context.Context, message string, fields ...zap.Field) {
	ctxFields := setFieldFromContext(ctx)
	logger.With(ctxFields...).Fatal(message, fields...)
}
