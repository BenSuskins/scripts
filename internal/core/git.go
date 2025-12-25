package core

import (
	"os/exec"
	"strings"
)

// OperationResult represents the result of a git operation
type OperationResult struct {
	Directory string
	Success   bool
	Message   string
}

// GetBranch returns the current branch for the given repo
func GetBranch(repoPath string) OperationResult {
	cmd := exec.Command("git", "-C", repoPath, "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return OperationResult{
			Directory: repoPath,
			Success:   false,
			Message:   strings.TrimSpace(string(output)),
		}
	}

	return OperationResult{
		Directory: repoPath,
		Success:   true,
		Message:   strings.TrimSpace(string(output)),
	}
}

// PullMain pulls from origin main
func PullMain(repoPath string) OperationResult {
	cmd := exec.Command("git", "-C", repoPath, "pull", "origin", "main")
	output, err := cmd.CombinedOutput()

	message := strings.TrimSpace(string(output))
	if message == "" {
		message = "Already up to date"
	}
	// Truncate long messages for compact display
	if len(message) > 60 {
		message = message[:57] + "..."
	}

	if err != nil {
		return OperationResult{
			Directory: repoPath,
			Success:   false,
			Message:   message,
		}
	}

	return OperationResult{
		Directory: repoPath,
		Success:   true,
		Message:   message,
	}
}
