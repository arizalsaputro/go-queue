package kq

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type contextKey struct{}

var headersContextKey = &contextKey{}

// WithHeaders injects Kafka headers into the context.
func WithHeaders(ctx context.Context, headers []kafka.Header) context.Context {
	return context.WithValue(ctx, headersContextKey, headers)
}

// GetHeaders retrieves Kafka headers from the context.
func GetHeaders(ctx context.Context) []kafka.Header {
	if headers, ok := ctx.Value(headersContextKey).([]kafka.Header); ok {
		return headers
	}
	return nil
}

// GetHeaderValue retrieves the value of a specific header by key.
func GetHeaderValue(ctx context.Context, key string) string {
	headers := GetHeaders(ctx)
	if headers == nil {
		return ""
	}

	for _, header := range headers {
		if header.Key == key {
			return string(header.Value)
		}
	}

	return ""
}
