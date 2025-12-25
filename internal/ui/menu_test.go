package ui

import (
	"testing"

	"suskins/scripts/internal/config"
)

func TestMenuItem(t *testing.T) {
	item := MenuItem{
		title:      "Test Command",
		desc:       "Test description",
		command:    "echo test",
		cmdType:    config.TypeSingle,
		showOutput: true,
	}

	t.Run("Title returns title", func(t *testing.T) {
		if item.Title() != "Test Command" {
			t.Errorf("expected 'Test Command', got '%s'", item.Title())
		}
	})

	t.Run("Description returns desc", func(t *testing.T) {
		if item.Description() != "Test description" {
			t.Errorf("expected 'Test description', got '%s'", item.Description())
		}
	})

	t.Run("FilterValue returns title", func(t *testing.T) {
		if item.FilterValue() != "Test Command" {
			t.Errorf("expected 'Test Command', got '%s'", item.FilterValue())
		}
	})

	t.Run("Command returns command", func(t *testing.T) {
		if item.Command() != "echo test" {
			t.Errorf("expected 'echo test', got '%s'", item.Command())
		}
	})

	t.Run("CmdType returns cmdType", func(t *testing.T) {
		if item.CmdType() != config.TypeSingle {
			t.Errorf("expected TypeSingle, got '%s'", item.CmdType())
		}
	})

	t.Run("ShowOutput returns showOutput", func(t *testing.T) {
		if !item.ShowOutput() {
			t.Error("expected ShowOutput to be true")
		}
	})
}

func TestNewMenuItems(t *testing.T) {
	cfg := &config.Config{
		Commands: []config.Command{
			{
				Name:        "Command 1",
				Description: "First command",
				Command:     "echo 1",
				Type:        config.TypeSingle,
				ShowOutput:  true,
			},
			{
				Name:        "Command 2",
				Description: "Second command",
				Command:     "echo 2",
				Type:        config.TypeGitDirs,
				ShowOutput:  false,
			},
		},
	}

	items := NewMenuItems(cfg)

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	item1 := items[0].(MenuItem)
	if item1.Title() != "Command 1" {
		t.Errorf("expected 'Command 1', got '%s'", item1.Title())
	}
	if item1.CmdType() != config.TypeSingle {
		t.Errorf("expected TypeSingle, got '%s'", item1.CmdType())
	}

	item2 := items[1].(MenuItem)
	if item2.Title() != "Command 2" {
		t.Errorf("expected 'Command 2', got '%s'", item2.Title())
	}
	if item2.CmdType() != config.TypeGitDirs {
		t.Errorf("expected TypeGitDirs, got '%s'", item2.CmdType())
	}
}
