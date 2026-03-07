package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveVault string `json:"active_vault"`
}

func PrepareConfig() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appDir := filepath.Join(configDir, "lockleaf")
	configPath := filepath.Join(appDir, "config.json")

	// Ensure Directory Exists
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Ensure File Exists (Create default if missing)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{ActiveVault: ""}
		data, _ := json.MarshalIndent(defaultConfig, "", "  ")
		return os.WriteFile(configPath, data, 0644)
	}

	return nil
}

func GetConfigPath() (string, error) {
	// Set config directory to %APPDATA% on Windows
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(configDir, "lockleaf")
	configPath := filepath.Join(appDir, "config.json")

	return configPath, nil
}

func SetActiveVault(name string) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	cfg := Config{ActiveVault: name}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func GetActiveVault() string {
	path, err := GetConfigPath()
	if err != nil {
		return ""
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg.ActiveVault
}
