package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type CommandType string

const (
	TypeSingle  CommandType = "single"
	TypeGitDirs CommandType = "git_dirs"
)

type Command struct {
	Name       string      `yaml:"name"`
	Category   string      `yaml:"category"`
	Command    string      `yaml:"command"`
	Type       CommandType `yaml:"type"`
	ShowOutput bool        `yaml:"show_output"`
}

type Config struct {
	Commands []Command `yaml:"commands"`
}

var defaultConfig = Config{
	Commands: []Command{
		{
			Name:       "Show git branch",
			Category:   "Git",
			Command:    "git rev-parse --abbrev-ref HEAD",
			Type:       TypeGitDirs,
			ShowOutput: true,
		},
		{
			Name:       "Pull main",
			Category:   "Git",
			Command:    "git pull origin main",
			Type:       TypeGitDirs,
			ShowOutput: false,
		},
	},
}

func configPaths() []string {
	paths := []string{}

	// Check ~/.config/scripts/config.yaml
	if home, err := os.UserHomeDir(); err == nil {
		paths = append(paths, filepath.Join(home, ".config", "scripts", "config.yaml"))
	}

	// Check ./scripts.yaml in current directory
	if cwd, err := os.Getwd(); err == nil {
		paths = append(paths, filepath.Join(cwd, "scripts.yaml"))
	}

	return paths
}

func LoadConfig() (*Config, error) {
	for _, path := range configPaths() {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var cfg Config
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", path, err)
		}

		return &cfg, nil
	}

	// No config found, create default
	if err := createDefaultConfig(); err != nil {
		// If we can't create, just return defaults in memory
		return &defaultConfig, nil
	}

	return &defaultConfig, nil
}

func createDefaultConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ".config", "scripts")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.yaml")

	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return err
	}

	header := "# Scripts TUI Configuration\n# Commands will appear in the menu\n#\n# type: \"single\" - runs once in current directory\n# type: \"git_dirs\" - runs in each git repository found in current directory\n\n"

	if err := os.WriteFile(configPath, []byte(header+string(data)), 0644); err != nil {
		return err
	}

	fmt.Printf("Created config at %s\n", configPath)
	return nil
}
