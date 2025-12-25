package core

import (
	"strings"
	"testing"
)

func TestRunCommand(t *testing.T) {
	t.Run("successful command", func(t *testing.T) {
		result := RunCommand(t.TempDir(), "echo hello")

		if !result.Success {
			t.Errorf("expected success, got failure: %s", result.Message)
		}
		if result.Message != "hello" {
			t.Errorf("expected 'hello', got '%s'", result.Message)
		}
	})

	t.Run("failed command", func(t *testing.T) {
		result := RunCommand(t.TempDir(), "exit 1")

		if result.Success {
			t.Error("expected failure, got success")
		}
	})

	t.Run("command with output on failure", func(t *testing.T) {
		result := RunCommand(t.TempDir(), "echo error message && exit 1")

		if result.Success {
			t.Error("expected failure, got success")
		}
		if !strings.Contains(result.Message, "error message") {
			t.Errorf("expected message to contain 'error message', got '%s'", result.Message)
		}
	})

	t.Run("command captures stderr", func(t *testing.T) {
		result := RunCommand(t.TempDir(), "echo stderr >&2")

		if !result.Success {
			t.Errorf("expected success, got failure: %s", result.Message)
		}
		if result.Message != "stderr" {
			t.Errorf("expected 'stderr', got '%s'", result.Message)
		}
	})

	t.Run("invalid directory", func(t *testing.T) {
		result := RunCommand("/nonexistent/path", "echo hello")

		if result.Success {
			t.Error("expected failure for invalid directory")
		}
		if result.Message == "" {
			t.Error("expected error message")
		}
	})

	t.Run("directory is set correctly", func(t *testing.T) {
		tmpDir := t.TempDir()
		result := RunCommand(tmpDir, "pwd")

		if !result.Success {
			t.Errorf("expected success, got failure: %s", result.Message)
		}
		if result.Message != tmpDir {
			t.Errorf("expected '%s', got '%s'", tmpDir, result.Message)
		}
		if result.Directory != tmpDir {
			t.Errorf("expected Directory to be '%s', got '%s'", tmpDir, result.Directory)
		}
	})
}
