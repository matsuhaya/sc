package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/matsumotohayato/sc/internal/config"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		cmdList()
		return
	}

	switch os.Args[1] {
	case "list":
		cmdList()
	case "new":
		cmdNew()
	case "edit":
		cmdEdit()
	case "help", "-h", "--help":
		cmdHelp()
	case "version", "-v", "--version":
		cmdVersion()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		fmt.Fprintln(os.Stderr, "Run 'sc help' for usage.")
		os.Exit(1)
	}
}

func cmdList() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if len(cfg.Shortcuts) == 0 {
		fmt.Println("No shortcuts registered.")
		fmt.Println("Use 'sc new <key> <description>' to add a shortcut.")
		return
	}

	maxKeyLen := 0
	for _, s := range cfg.Shortcuts {
		if len(s.Key) > maxKeyLen {
			maxKeyLen = len(s.Key)
		}
	}

	for _, s := range cfg.Shortcuts {
		padding := strings.Repeat(" ", maxKeyLen-len(s.Key))
		fmt.Printf("%s%s  %s\n", s.Key, padding, s.Description)
	}
}

func cmdNew() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage: sc new <key> <description>")
		fmt.Fprintln(os.Stderr, "Example: sc new \"Cmd+Shift+P\" \"Command Palette\"")
		os.Exit(1)
	}

	key := os.Args[2]
	description := os.Args[3]

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	cfg.Add(key, description)

	if err := config.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Added: %s  %s\n", key, description)
}

func cmdEdit() {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	path, err := config.ConfigPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting config path: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfg := &config.Config{Shortcuts: []config.Shortcut{}}
		if err := config.Save(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating config file: %v\n", err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("sh", "-c", editor+" "+path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running editor: %v\n", err)
		os.Exit(1)
	}
}

func cmdHelp() {
	help := `sc - VS Code Shortcut Manager

Usage:
  sc              Show all shortcuts (same as 'sc list')
  sc list         Show all shortcuts
  sc new <key> <description>
                  Add a new shortcut
  sc edit         Open config file in editor
  sc help         Show this help
  sc version      Show version

Examples:
  sc new "Cmd+Shift+P" "Command Palette"
  sc new "Cmd+P" "Quick Open"

Config file: ~/.config/sc/shortcuts.yaml`

	fmt.Println(help)
}

func cmdVersion() {
	fmt.Printf("sc version %s\n", version)
}
