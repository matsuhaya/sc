package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Shortcut struct {
	Key         string `yaml:"key"`
	Description string `yaml:"description"`
}

type Config struct {
	Shortcuts []Shortcut `yaml:"shortcuts"`
}

func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "sc"), nil
}

func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "shortcuts.yaml"), nil
}

func ensureConfigDir() error {
	dir, err := ConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(dir, 0755)
}

func Load() (*Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Shortcuts: []Shortcut{}}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	path, err := ConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (c *Config) Add(key, description string) {
	c.Shortcuts = append(c.Shortcuts, Shortcut{
		Key:         key,
		Description: description,
	})
}
