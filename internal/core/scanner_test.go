package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGitDirectories(t *testing.T) {
	t.Run("empty directory returns empty slice", func(t *testing.T) {
		tmpDir := t.TempDir()

		dirs, err := FindGitDirectories(tmpDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(dirs) != 0 {
			t.Errorf("expected empty slice, got %v", dirs)
		}
	})

	t.Run("finds git repositories", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create two git repos
		repo1 := filepath.Join(tmpDir, "repo-a")
		repo2 := filepath.Join(tmpDir, "repo-b")
		os.Mkdir(repo1, 0755)
		os.Mkdir(repo2, 0755)
		os.Mkdir(filepath.Join(repo1, ".git"), 0755)
		os.Mkdir(filepath.Join(repo2, ".git"), 0755)

		dirs, err := FindGitDirectories(tmpDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(dirs) != 2 {
			t.Fatalf("expected 2 repos, got %d", len(dirs))
		}
		// Should be sorted
		if filepath.Base(dirs[0]) != "repo-a" {
			t.Errorf("expected repo-a first, got %s", filepath.Base(dirs[0]))
		}
		if filepath.Base(dirs[1]) != "repo-b" {
			t.Errorf("expected repo-b second, got %s", filepath.Base(dirs[1]))
		}
	})

	t.Run("ignores non-git directories", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create one git repo and one regular dir
		gitRepo := filepath.Join(tmpDir, "git-repo")
		regularDir := filepath.Join(tmpDir, "regular-dir")
		os.Mkdir(gitRepo, 0755)
		os.Mkdir(regularDir, 0755)
		os.Mkdir(filepath.Join(gitRepo, ".git"), 0755)

		dirs, err := FindGitDirectories(tmpDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(dirs) != 1 {
			t.Fatalf("expected 1 repo, got %d", len(dirs))
		}
		if filepath.Base(dirs[0]) != "git-repo" {
			t.Errorf("expected git-repo, got %s", filepath.Base(dirs[0]))
		}
	})

	t.Run("ignores files", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create a file (not a directory)
		os.WriteFile(filepath.Join(tmpDir, "somefile.txt"), []byte("hello"), 0644)

		dirs, err := FindGitDirectories(tmpDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(dirs) != 0 {
			t.Errorf("expected empty slice, got %v", dirs)
		}
	})

	t.Run("non-existent directory returns error", func(t *testing.T) {
		_, err := FindGitDirectories("/nonexistent/path/that/does/not/exist")
		if err == nil {
			t.Error("expected error for non-existent directory")
		}
	})
}
