package helper

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"go-test-api/internal/model"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	traceKey        = "logging.googleapis.com/trace"
	spanKey         = "logging.googleapis.com/spanId"
	traceSampledKey = "logging.googleapis.com/trace_sampled"
	httpLog         = "[HTTP]"
)

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func ZapLogger(ctx context.Context, conf model.Config, log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			traceHeader := req.Header.Get("X-Cloud-Trace-Context")
			traceID, spanID, traceSampled := deconstructXCloudTraceContext(traceHeader)
			traceID = fmt.Sprintf("projects/traces/%s", traceID)

			fields := []zapcore.Field{
				zap.String("traceparent", req.Header.Get("Traceparent")),
				zap.String(traceKey, traceID),
				zap.String(spanKey, spanID),
				zap.Bool(traceSampledKey, traceSampled),
				zap.String("referenceNumber", req.Header.Get("referenceNumber")),
				zap.String("host", req.Host),
				zap.String("real_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
				zap.String("pid", strconv.Itoa(os.Getpid())),
				zap.String("path", req.URL.String()),
				zap.String("method", req.Method),
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = uuid.New().String()
			}
			fields = append(fields, zap.String("request_id", id))

			n := res.Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error(httpLog, fields...)
			case n >= 400:
				log.With(zap.Error(err)).Warn(httpLog, fields...)
			case n >= 300:
				log.Info(httpLog, fields...)
			default:
				log.Info(httpLog, fields...)
			}

			return nil
		}
	}
}
