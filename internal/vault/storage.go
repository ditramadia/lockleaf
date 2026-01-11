package vault

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Storage struct {
	DataDir string
}

// NewStorage creates a new storage instance with the given path
func NewStorage(dataDir string) *Storage {
	return &Storage{DataDir: dataDir}
}

// GetPath returns the full file path for the given vault name
func (s *Storage) GetPath(vaultName string) string {
	return filepath.Join(s.DataDir, vaultName+".leaf")
}

func (s *Storage) Save(v *Vault) error {
	path := s.GetPath(v.Name)

	// Marshal the data
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("[ERROR] failed to encode vault: %w", err)
	}

	// Create a temporary file
	tmpPath := path + ".tmp"
	tmpFile, err := os.OpenFile(tmpPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	// Write the data to temp file
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("[ERROR] failed to write to temp file: %w", err)
	}

	// Force sync to physical disk
	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("[ERROR] failed to sync file: %w", err)
	}

	// Atomic rename (replaces old file with new one)
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("[ERROR] failed to finalize save: %w", err)
	}

	return nil
}

func (s *Storage) Load(vaultName string) (*Vault, error) {
	path := s.GetPath(vaultName)

	// Open the vault file
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("[ERROR] vault '%s' not found", vaultName)
		}
		return nil, err
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to read vault file: %w", err)
	}

	// Unmarshal the JSON data
	var v Vault
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("[ERROR] failed to parse vault JSON: %w", err)
	}

	return &v, nil
}
