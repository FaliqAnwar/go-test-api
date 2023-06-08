package ctxdata

import "context"

type traceParent struct{}
type traceID struct{}
type spanID struct{}
type traceSampled struct{}
type referenceNumber struct{}
type host struct{}
type realIP struct{}
type userAgent struct{}
type pid struct{}
type path struct{}
type method struct{}
type status struct{}

type Set func(ctx context.Context) context.Context

func Sets(ctx context.Context, fn ...Set) context.Context {
	for _, f := range fn {
		ctx = f(ctx)
	}
	return ctx
}

func SetTraceParent(tp string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceParent{}, tp)
	}
}

func SetTraceID(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceID{}, s)
	}
}

func SetSpanID(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, spanID{}, s)
	}
}

func SetTraceSampled(b bool) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, traceSampled{}, b)
	}
}

func SetReferenceNumber(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, referenceNumber{}, s)
	}
}

func SetHost(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, host{}, s)
	}
}

func SetRealIP(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, realIP{}, s)
	}
}

func SetUserAgent(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userAgent{}, s)
	}
}

func SetPid(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pid{}, s)
	}
}

func SetPath(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, path{}, s)
	}
}

func SetMethod(s string) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, method{}, s)
	}
}

func SetStatus(i int) Set {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, status{}, i)
	}
}

func GetTraceParent(ctx context.Context) string {
	if v, ok := ctx.Value(traceParent{}).(string); ok {
		return v
	}
	return ""
}

func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.Value(traceID{}).(string); ok {
		return v
	}
	return ""
}

func GetSpanID(ctx context.Context) string {
	if v, ok := ctx.Value(spanID{}).(string); ok {
		return v
	}
	return ""
}

func GetTraceSampled(ctx context.Context) bool {
	if v, ok := ctx.Value(traceSampled{}).(bool); ok {
		return v
	}
	return false
}

func GetReferenceNumber(ctx context.Context) string {
	if v, ok := ctx.Value(referenceNumber{}).(string); ok {
		return v
	}
	return ""
}

func GetHost(ctx context.Context) string {
	if v, ok := ctx.Value(host{}).(string); ok {
		return v
	}
	return ""
}

func GetRealIP(ctx context.Context) string {
	if v, ok := ctx.Value(realIP{}).(string); ok {
		return v
	}
	return ""
}

func GetUserAgent(ctx context.Context) string {
	if v, ok := ctx.Value(userAgent{}).(string); ok {
		return v
	}
	return ""
}

func GetPid(ctx context.Context) string {
	if v, ok := ctx.Value(pid{}).(string); ok {
		return v
	}
	return ""
}

func GetPath(ctx context.Context) string {
	if v, ok := ctx.Value(path{}).(string); ok {
		return v
	}
	return ""
}

func GetMethod(ctx context.Context) string {
	if v, ok := ctx.Value(method{}).(string); ok {
		return v
	}
	return ""
}

func GetStatus(ctx context.Context) int {
	if v, ok := ctx.Value(status{}).(int); ok {
		return v
	}
	return 0
}
