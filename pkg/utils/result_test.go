package utils

import (
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewToolResultText(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "simple message",
			message: "Success",
		},
		{
			name:    "empty message",
			message: "",
		},
		{
			name:    "long message",
			message: "This is a very long message that contains multiple sentences and should still be handled correctly by the function.",
		},
		{
			name:    "message with newlines",
			message: "Line 1\nLine 2\nLine 3",
		},
		{
			name:    "message with special characters",
			message: "Special chars: !@#$%^&*()_+-=[]{}|;':\",./<>?",
		},
		{
			name:    "unicode message",
			message: "Hello 世界 🌍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewToolResultText(tt.message)

			require.NotNil(t, result)
			assert.False(t, result.IsError)
			require.Len(t, result.Content, 1)

			textContent, ok := result.Content[0].(*mcp.TextContent)
			require.True(t, ok, "expected TextContent")
			assert.Equal(t, tt.message, textContent.Text)
		})
	}
}

func TestNewToolResultError(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "simple error message",
			message: "Something went wrong",
		},
		{
			name:    "empty error message",
			message: "",
		},
		{
			name:    "error with details",
			message: "Failed to process request: invalid input",
		},
		{
			name:    "error with newlines",
			message: "Error:\n  - First issue\n  - Second issue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewToolResultError(tt.message)

			require.NotNil(t, result)
			assert.True(t, result.IsError, "IsError should be true")
			require.Len(t, result.Content, 1)

			textContent, ok := result.Content[0].(*mcp.TextContent)
			require.True(t, ok, "expected TextContent")
			assert.Equal(t, tt.message, textContent.Text)
		})
	}
}

func TestNewToolResultErrorFromErr(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		err         error
		wantText    string
	}{
		{
			name:     "error with standard error",
			message:  "Operation failed",
			err:      errors.New("connection timeout"),
			wantText: "Operation failed: connection timeout",
		},
		{
			name:     "error with wrapped error",
			message:  "Failed to save",
			err:      errors.New("disk full"),
			wantText: "Failed to save: disk full",
		},
		{
			name:     "empty message with error",
			message:  "",
			err:      errors.New("some error"),
			wantText: ": some error",
		},
		{
			name:     "message with error containing special chars",
			message:  "Error",
			err:      errors.New("failed: %v"),
			wantText: "Error: failed: %v",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewToolResultErrorFromErr(tt.message, tt.err)

			require.NotNil(t, result)
			assert.True(t, result.IsError, "IsError should be true")
			require.Len(t, result.Content, 1)

			textContent, ok := result.Content[0].(*mcp.TextContent)
			require.True(t, ok, "expected TextContent")
			assert.Equal(t, tt.wantText, textContent.Text)
		})
	}
}

func TestNewToolResultResource(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		contents *mcp.ResourceContents
	}{
		{
			name:    "resource with text content",
			message: "File contents",
			contents: &mcp.ResourceContents{
				URI:      "file:///path/to/file.txt",
				MIMEType: "text/plain",
				Text:     "Hello, World!",
			},
		},
		{
			name:    "resource with empty message",
			message: "",
			contents: &mcp.ResourceContents{
				URI:  "file:///path/to/data.json",
				Text: `{"key": "value"}`,
			},
		},
		{
			name:    "resource with blob content",
			message: "Binary data",
			contents: &mcp.ResourceContents{
				URI:      "file:///path/to/image.png",
				MIMEType: "image/png",
				Blob:     []byte("base64encodeddata"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewToolResultResource(tt.message, tt.contents)

			require.NotNil(t, result)
			assert.False(t, result.IsError, "IsError should be false")
			require.Len(t, result.Content, 2, "should have text and embedded resource")

			// Check text content
			textContent, ok := result.Content[0].(*mcp.TextContent)
			require.True(t, ok, "first content should be TextContent")
			assert.Equal(t, tt.message, textContent.Text)

			// Check embedded resource
			embeddedResource, ok := result.Content[1].(*mcp.EmbeddedResource)
			require.True(t, ok, "second content should be EmbeddedResource")
			assert.Equal(t, tt.contents, embeddedResource.Resource)
		})
	}
}

func TestNewToolResultTypes(t *testing.T) {
	t.Run("text result should not be error", func(t *testing.T) {
		result := NewToolResultText("test")
		assert.False(t, result.IsError)
	})

	t.Run("error result should be error", func(t *testing.T) {
		result := NewToolResultError("test")
		assert.True(t, result.IsError)
	})

	t.Run("error from err result should be error", func(t *testing.T) {
		result := NewToolResultErrorFromErr("test", errors.New("err"))
		assert.True(t, result.IsError)
	})

	t.Run("resource result should not be error", func(t *testing.T) {
		result := NewToolResultResource("test", &mcp.ResourceContents{URI: "test"})
		assert.False(t, result.IsError)
	})
}

// TODO: Add benchmark tests if these functions are used in hot paths
// func BenchmarkNewToolResultText(b *testing.B) { ... }
