package context

import (
	"context"
	"crypto/rand"
	"fmt"
)

const correlationIDKey contextKey = "correlationID"

func NewCorrelationId() string {
	buf := make([]byte, 4)
	rand.Read(buf)
	return fmt.Sprintf("id-%x", buf)
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	if correlationID == "" {
		correlationID = GetCorrelationID(ctx)
		if correlationID == "" {
			correlationID = NewCorrelationId()
		}
	}
	return context.WithValue(ctx, correlationIDKey, correlationID)
}

func GetCorrelationID(ctx context.Context) string {
	if v, ok := ctx.Value(correlationIDKey).(string); ok {
		return v
	}
	return ""
}
