package requestid

import "context"

type contextKey struct{}

var requestIDKey = contextKey{}

func SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func GetRequestIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(requestIDKey)
	id, ok := v.(string)
	return id, ok
}
