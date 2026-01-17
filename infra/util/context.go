package util

import "context"

type CtxField string

const (
	CtxTraceId CtxField = "X-Trace-Id"
	CtxSpanId  CtxField = "X-Span-Id"
	CtxReqBody CtxField = "X-Req-Body"
	CtxResBody CtxField = "X-Res-Body"
)

type Pair[K, V any] struct {
	Key K
	Val V
}

var CtxPairs = []Pair[CtxField, string]{
	{CtxTraceId, "traceId="},
	{CtxSpanId, "spanId="},
	{CtxReqBody, "req="},
	{CtxResBody, "res="},
}

func CtxString(ctx context.Context, k CtxField) string {
	if ctx == nil {
		return ""
	}
	if val, _ := ctx.Value(k).(string); val != "" {
		return val
	}

	return ""
}

func CtxWith(ctx context.Context, k CtxField, v string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, k, v)
}
