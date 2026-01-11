package vault

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// Setup temporary directory for the test
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	// Create a dummy vault
	v := NewVault("test-vault")
	c := Credential{
		Name: "Github",
		Fields: []Field{
			{Label: "username", Value: "gopher", IsSecret: false},
			{Label: "password", Value: "s3cr3t", IsSecret: true},
		},
	}
	v.Credentials = append(v.Credentials, c)

	// Test saving the vault
	err := storage.Save(v)
	if err != nil {
		t.Fatalf("Failed to save vault: %v", err)
	}

	// Verify file exists physically
	expectedPath := filepath.Join(tmpDir, "test-vault.leaf")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Vault file was not created at %s", expectedPath)
	}

	// Test loading the vault
	loadedVault, err := storage.Load("test-vault")
	if err != nil {
		t.Fatalf("Failed to load vault: %v", err)
	}

	// Assert data integrity
	if loadedVault.Name != v.Name {
		t.Errorf("Expected name %s, got %s", v.Name, loadedVault.Name)
	}

	if c := loadedVault.Credentials[0]; c.Name != "Github" {
		t.Errorf("Missing credential 'github' after load")
	}
}

func TestLoadMissingVault(t *testing.T) {
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	_, err := storage.Load("non-existent")
	if err == nil {
		t.Errorf("Expected error when loading non-existent vault, got nil")
	}
}
