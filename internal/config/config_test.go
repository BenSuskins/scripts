package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfigParsing(t *testing.T) {
	t.Run("valid YAML parses correctly", func(t *testing.T) {
		yamlData := `
commands:
  - name: "Test Command"
    description: "A test command"
    command: "echo hello"
    type: "single"
    show_output: true
  - name: "Git Command"
    description: "Run in git dirs"
    command: "git status"
    type: "git_dirs"
    show_output: false
`
		var cfg Config
		err := yaml.Unmarshal([]byte(yamlData), &cfg)
		if err != nil {
			t.Fatalf("failed to parse YAML: %v", err)
		}

		if len(cfg.Commands) != 2 {
			t.Fatalf("expected 2 commands, got %d", len(cfg.Commands))
		}

		cmd1 := cfg.Commands[0]
		if cmd1.Name != "Test Command" {
			t.Errorf("expected name 'Test Command', got '%s'", cmd1.Name)
		}
		if cmd1.Type != TypeSingle {
			t.Errorf("expected type 'single', got '%s'", cmd1.Type)
		}
		if !cmd1.ShowOutput {
			t.Error("expected ShowOutput to be true")
		}

		cmd2 := cfg.Commands[1]
		if cmd2.Type != TypeGitDirs {
			t.Errorf("expected type 'git_dirs', got '%s'", cmd2.Type)
		}
		if cmd2.ShowOutput {
			t.Error("expected ShowOutput to be false")
		}
	})

	t.Run("invalid YAML returns error", func(t *testing.T) {
		yamlData := `
commands:
  - name: [invalid yaml structure
`
		var cfg Config
		err := yaml.Unmarshal([]byte(yamlData), &cfg)
		if err == nil {
			t.Error("expected error for invalid YAML")
		}
	})

	t.Run("empty YAML returns empty commands", func(t *testing.T) {
		yamlData := `commands: []`

		var cfg Config
		err := yaml.Unmarshal([]byte(yamlData), &cfg)
		if err != nil {
			t.Fatalf("failed to parse YAML: %v", err)
		}

		if len(cfg.Commands) != 0 {
			t.Errorf("expected 0 commands, got %d", len(cfg.Commands))
		}
	})
}

func TestCommandType(t *testing.T) {
	t.Run("TypeSingle constant", func(t *testing.T) {
		if TypeSingle != "single" {
			t.Errorf("expected 'single', got '%s'", TypeSingle)
		}
	})

	t.Run("TypeGitDirs constant", func(t *testing.T) {
		if TypeGitDirs != "git_dirs" {
			t.Errorf("expected 'git_dirs', got '%s'", TypeGitDirs)
		}
	})
}

func TestDefaultConfig(t *testing.T) {
	// defaultConfig should have at least one command
	if len(defaultConfig.Commands) == 0 {
		t.Error("defaultConfig should have at least one command")
	}

	// First command should be "Branch"
	if defaultConfig.Commands[0].Name != "Branch" {
		t.Errorf("expected first command to be 'Branch', got '%s'", defaultConfig.Commands[0].Name)
	}
}
