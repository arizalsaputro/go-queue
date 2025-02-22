package kq

import (
	"context"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

// TestWithHeaders tests the WithHeaders function.
func TestWithHeaders(t *testing.T) {
	// Create sample headers
	headers := []kafka.Header{
		{Key: "Trace-ID", Value: []byte("12345")},
		{Key: "User-ID", Value: []byte("67890")},
	}

	// Inject headers into context
	ctx := WithHeaders(context.Background(), headers)

	// Retrieve headers from context
	retrievedHeaders := GetHeaders(ctx)

	// Assert that the retrieved headers match the original headers
	assert.Equal(t, headers, retrievedHeaders, "Retrieved headers should match the injected headers")
}

// TestGetHeaders_NoHeaders tests GetHeaders when no headers are present in the context.
func TestGetHeaders_NoHeaders(t *testing.T) {
	// Create a context without headers
	ctx := context.Background()

	// Retrieve headers from context
	retrievedHeaders := GetHeaders(ctx)

	// Assert that no headers are retrieved
	assert.Nil(t, retrievedHeaders, "No headers should be retrieved from the context")
}

// TestGetHeaderValue tests the GetHeaderValue function.
func TestGetHeaderValue(t *testing.T) {
	// Create sample headers
	headers := []kafka.Header{
		{Key: "Trace-ID", Value: []byte("12345")},
		{Key: "User-ID", Value: []byte("67890")},
	}

	// Inject headers into context
	ctx := WithHeaders(context.Background(), headers)

	// Test retrieving an existing header
	traceID := GetHeaderValue(ctx, "Trace-ID")
	assert.Equal(t, "12345", traceID, "Trace-ID should be '12345'")

	// Test retrieving another existing header
	userID := GetHeaderValue(ctx, "User-ID")
	assert.Equal(t, "67890", userID, "User-ID should be '67890'")

	// Test retrieving a non-existent header
	nonExistent := GetHeaderValue(ctx, "Non-Existent-Key")
	assert.Equal(t, "", nonExistent, "Non-existent key should return an empty string")
}

// TestGetHeaderValue_NoHeaders tests GetHeaderValue when no headers are present in the context.
func TestGetHeaderValue_NoHeaders(t *testing.T) {
	// Create a context without headers
	ctx := context.Background()

	// Test retrieving a header from a context without headers
	value := GetHeaderValue(ctx, "Trace-ID")
	assert.Equal(t, "", value, "Value should be empty when no headers are present in the context")
}
