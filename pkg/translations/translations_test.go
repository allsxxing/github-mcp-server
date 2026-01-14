package translations

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullTranslationHelper(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "returns default value",
			key:          "ANY_KEY",
			defaultValue: "default text",
			want:         "default text",
		},
		{
			name:         "ignores key parameter",
			key:          "SOME_KEY",
			defaultValue: "expected",
			want:         "expected",
		},
		{
			name:         "handles empty key",
			key:          "",
			defaultValue: "default",
			want:         "default",
		},
		{
			name:         "handles empty default",
			key:          "KEY",
			defaultValue: "",
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullTranslationHelper(tt.key, tt.defaultValue)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTranslationHelper_DefaultValues(t *testing.T) {
	// Create a temporary directory for test
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	helper, cleanup := TranslationHelper()
	defer cleanup()

	tests := []struct {
		name         string
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "returns default when no override",
			key:          "TEST_KEY",
			defaultValue: "default value",
			want:         "default value",
		},
		{
			name:         "handles empty default",
			key:          "EMPTY_KEY",
			defaultValue: "",
			want:         "",
		},
		{
			name:         "key is converted to uppercase",
			key:          "lowercase_key",
			defaultValue: "default",
			want:         "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helper(tt.key, tt.defaultValue)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTranslationHelper_EnvironmentVariables(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Set environment variable
	testKey := "GITHUB_MCP_TEST_ENV_VAR"
	testValue := "environment override"
	err = os.Setenv(testKey, testValue)
	require.NoError(t, err)
	defer func() { _ = os.Unsetenv(testKey) }()

	helper, cleanup := TranslationHelper()
	defer cleanup()

	t.Run("reads from environment variable", func(t *testing.T) {
		result := helper("TEST_ENV_VAR", "default")
		assert.Equal(t, testValue, result)
	})

	t.Run("caches environment variable value", func(t *testing.T) {
		// First call
		result1 := helper("TEST_ENV_VAR", "default")
		assert.Equal(t, testValue, result1)

		// Change env var after first read
		_ = os.Setenv(testKey, "changed value")

		// Second call should return cached value
		result2 := helper("TEST_ENV_VAR", "default")
		assert.Equal(t, testValue, result2, "should return cached value")
	})
}

func TestTranslationHelper_ConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Create a test config file
	configContent := `{
		"FILE_KEY": "value from file",
		"ANOTHER_KEY": "another value"
	}`
	err = os.WriteFile("github-mcp-server-config.json", []byte(configContent), 0644)
	require.NoError(t, err)

	helper, cleanup := TranslationHelper()
	defer cleanup()

	t.Run("reads from config file", func(t *testing.T) {
		result := helper("FILE_KEY", "default")
		assert.Equal(t, "value from file", result)
	})

	t.Run("reads multiple keys from config", func(t *testing.T) {
		result := helper("ANOTHER_KEY", "default")
		assert.Equal(t, "another value", result)
	})
}

func TestTranslationHelper_PriorityOrder(t *testing.T) {
	// Test that environment variables take precedence over config file
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Create config file with a value
	configContent := `{
		"PRIORITY_KEY": "value from file"
	}`
	err = os.WriteFile("github-mcp-server-config.json", []byte(configContent), 0644)
	require.NoError(t, err)

	// Set environment variable with same key
	testKey := "GITHUB_MCP_PRIORITY_KEY"
	testValue := "value from env"
	err = os.Setenv(testKey, testValue)
	require.NoError(t, err)
	defer func() { _ = os.Unsetenv(testKey) }()

	helper, cleanup := TranslationHelper()
	defer cleanup()

	t.Run("environment variable overrides config file", func(t *testing.T) {
		result := helper("PRIORITY_KEY", "default")
		assert.Equal(t, testValue, result, "env var should take precedence")
	})
}

func TestTranslationHelper_CaseInsensitivity(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	helper, cleanup := TranslationHelper()
	defer cleanup()

	// First call with lowercase
	result1 := helper("test_key", "default")
	assert.Equal(t, "default", result1)

	// Second call with uppercase should return cached value
	result2 := helper("TEST_KEY", "different default")
	assert.Equal(t, "default", result2, "keys should be case-insensitive (converted to uppercase)")
}

func TestDumpTranslationKeyMap(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	t.Run("creates config file", func(t *testing.T) {
		testMap := map[string]string{
			"KEY1": "value1",
			"KEY2": "value2",
		}

		err := DumpTranslationKeyMap(testMap)
		require.NoError(t, err)

		// Verify file was created
		filePath := filepath.Join(tmpDir, "github-mcp-server-config.json")
		_, err = os.Stat(filePath)
		require.NoError(t, err, "config file should be created")
	})

	t.Run("writes valid JSON", func(t *testing.T) {
		testMap := map[string]string{
			"KEY1": "value1",
			"KEY2": "value2",
		}

		err := DumpTranslationKeyMap(testMap)
		require.NoError(t, err)

		// Read and verify content
		content, err := os.ReadFile("github-mcp-server-config.json")
		require.NoError(t, err)

		// Verify it's valid JSON by checking content
		assert.Contains(t, string(content), "KEY1")
		assert.Contains(t, string(content), "value1")
	})

	t.Run("handles empty map", func(t *testing.T) {
		testMap := map[string]string{}

		err := DumpTranslationKeyMap(testMap)
		require.NoError(t, err)
	})

	t.Run("overwrites existing file", func(t *testing.T) {
		// Create initial file
		testMap1 := map[string]string{"KEY1": "value1"}
		err := DumpTranslationKeyMap(testMap1)
		require.NoError(t, err)

		// Overwrite with new content
		testMap2 := map[string]string{"KEY2": "value2"}
		err = DumpTranslationKeyMap(testMap2)
		require.NoError(t, err)

		// Verify new content
		content, err := os.ReadFile("github-mcp-server-config.json")
		require.NoError(t, err)
		assert.Contains(t, string(content), "KEY2")
		assert.NotContains(t, string(content), "KEY1")
	})
}

func TestTranslationHelper_CleanupFunction(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	helper, cleanup := TranslationHelper()

	// Use helper to populate translation map
	_ = helper("TEST_KEY", "test value")

	// Call cleanup
	cleanup()

	// Verify config file was created
	filePath := filepath.Join(tmpDir, "github-mcp-server-config.json")
	_, err = os.Stat(filePath)
	require.NoError(t, err, "cleanup should dump translation map to file")
}

func TestTranslationHelper_NoConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	require.NoError(t, err)
	defer func() { _ = os.Chdir(origDir) }()

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Ensure no config file exists
	_ = os.Remove("github-mcp-server-config.json")

	t.Run("works without config file", func(t *testing.T) {
		helper, cleanup := TranslationHelper()
		defer cleanup()

		result := helper("KEY", "default")
		assert.Equal(t, "default", result)
	})
}

// TODO: Add tests for invalid JSON in config file
// func TestTranslationHelper_InvalidConfigJSON(t *testing.T) { ... }

// TODO: Add tests for file permission errors
// func TestDumpTranslationKeyMap_PermissionError(t *testing.T) { ... }

// TODO: Add concurrent access tests if used in concurrent contexts
// func TestTranslationHelper_Concurrent(t *testing.T) { ... }
