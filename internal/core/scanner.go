package core

import (
	"os"
	"path/filepath"
	"sort"
)

// FindGitDirectories walks the directory tree starting from root
// and returns all directories containing a .git folder.
func FindGitDirectories(root string) ([]string, error) {
	var gitDirs []string

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		path := filepath.Join(root, entry.Name())
		gitPath := filepath.Join(path, ".git")

		if info, err := os.Stat(gitPath); err == nil && info.IsDir() {
			gitDirs = append(gitDirs, path)
		}
	}

	sort.Strings(gitDirs)
	return gitDirs, nil
}
