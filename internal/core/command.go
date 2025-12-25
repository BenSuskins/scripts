package core

import (
	"os/exec"
	"strings"
)

// OperationResult represents the result of a command operation
type OperationResult struct {
	Directory string
	Success   bool
	Message   string
}

// RunCommand executes a shell command in the specified directory
func RunCommand(dir, command string) OperationResult {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	message := strings.TrimSpace(string(output))

	if err != nil {
		if message == "" {
			message = err.Error()
		}
		return OperationResult{
			Directory: dir,
			Success:   false,
			Message:   message,
		}
	}

	return OperationResult{
		Directory: dir,
		Success:   true,
		Message:   message,
	}
}
