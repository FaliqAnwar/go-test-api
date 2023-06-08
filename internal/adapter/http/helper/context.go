package helper

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"go-test-api/internal/ctxdata"
	"go-test-api/internal/model"

	echo "github.com/labstack/echo/v4"
)

func SetContext(conf model.Config) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			traceHeader := req.Header.Get("X-Cloud-Trace-Context")
			traceID, spanID, traceSampled := deconstructXCloudTraceContext(traceHeader)
			traceID = fmt.Sprintf("projects/traces/%s", traceID)

			cdata := ctxdata.Sets(
				c.Request().Context(),
				ctxdata.SetTraceParent(req.Header.Get("Traceparent")),
				ctxdata.SetTraceID(traceID),
				ctxdata.SetSpanID(spanID),
				ctxdata.SetTraceSampled(traceSampled),
				ctxdata.SetReferenceNumber(req.Header.Get("referenceNumber")),
				ctxdata.SetHost(req.Host),
				ctxdata.SetRealIP(c.RealIP()),
				ctxdata.SetUserAgent(req.UserAgent()),
				ctxdata.SetPid(strconv.Itoa(os.Getpid())),
				ctxdata.SetPath(req.URL.String()),
				ctxdata.SetMethod(req.Method),
				ctxdata.SetStatus(res.Status),
			)
			c.SetRequest(c.Request().WithContext(cdata))

			return next(c)
		}
	}
}

// taken from https://github.com/googleapis/google-cloud-go/blob/master/logging/logging.go#L774
var reCloudTraceContext = regexp.MustCompile(
	// Matches on "TRACE_ID"
	`([a-f\d]+)?` +
		// Matches on "/SPAN_ID"
		`(?:/([a-f\d]+))?` +
		// Matches on ";0=TRACE_TRUE"
		`(?:;o=(\d))?`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string, traceSampled bool) {
	// As per the format described at https://cloud.google.com/trace/docs/setup#force-trace
	//    "X-Cloud-Trace-Context: TRACE_ID/SPAN_ID;o=TRACE_TRUE"
	// for example:
	//    "X-Cloud-Trace-Context: 105445aa7843bc8bf206b120001000/1;o=1"
	//
	// We expect:
	//   * traceID (optional):          "105445aa7843bc8bf206b120001000"
	//   * spanID (optional):           "1"
	//   * traceSampled (optional):     true
	matches := reCloudTraceContext.FindStringSubmatch(s)
	traceID, spanID, traceSampled = matches[1], matches[2], matches[3] == "1"
	if spanID == "0" {
		spanID = ""
	}
	return
}
