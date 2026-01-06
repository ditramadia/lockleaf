package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	FieldEmail    FieldType = "email"
	FieldPhone    FieldType = "phone"
	FieldUsername FieldType = "username"
	FieldPassword FieldType = "password"
	FieldPIN      FieldType = "pin"
)

type FieldType string

type Field struct {
	Key   string    `json:"key"`
	Type  FieldType `json:"type"`
	Value string    `json:"value"`
}

type Credential struct {
	ID        string    `json:"id"`
	Label     string    `json:"label"`
	Category  string    `json:"category"`
	Fields    []Field   `json:"fields"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vault struct {
	Name        string       `json:"name"`
	Credentials []Credential `json:"credentials"`
}

func VaultExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func SaveVault(v Vault, path string) error {
	// Derive the save directory from the path
	var dir string = filepath.Dir(path)

	// Create save directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", dir, err)
	}

	// Marshal the vault and save as JSON (for now)
	data, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return fmt.Errorf("Failed to marshal vault: %w", err)
	}

	// Write file with restricted permission (0600)
	if err = os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("Failed to write vault file: %w", err)
	}

	return nil
}

func LoadVault(path string) (Vault, error) {
	var v Vault

	// Read the vault file
	data, err := os.ReadFile(path)
	if err != nil {
		return v, fmt.Errorf("Failed to read vault: %w", err)
	}

	// Unmarshal the vault
	if err = json.Unmarshal(data, &v); err != nil {
		return v, fmt.Errorf("Failed to unmarshal vault: %w", err)
	}

	return v, err
}
