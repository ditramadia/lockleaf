package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveVault string `json:"active_vault"`
	path        string
}

func Init() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(configDir, "lockleaf")
	configPath := filepath.Join(appDir, "config.json")

	// Ensure Directory Exists
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	// Ensure File Exists (Create default if missing)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cfg := &Config{
			ActiveVault: "",
			path:        configPath,
		}

		data, _ := json.MarshalIndent(cfg, "", "  ")
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return nil, err
		}
	}

	cfg := &Config{path: configPath}
	if err := cfg.Load(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) Load() error {
	data, err := os.ReadFile(c.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, data, 0644)
}

func (c *Config) GetConfigPath() (string, error) {
	return c.path, nil
}

func (c *Config) SetActiveVault(name string) error {
	c.ActiveVault = name
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.path, data, 0644)
}

func (c *Config) GetActiveVault() (string, error) {
	return c.ActiveVault, nil
}
