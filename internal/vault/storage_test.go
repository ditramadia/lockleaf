package vault

import (
	"os"
	"path/filepath"
	"testing"
)

// TestSaveAndLoad tests the saving and loading of vaults
func TestSaveAndLoad(t *testing.T) {
	// Setup temporary storage for the test
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	// Create a dummy vault
	v := NewVault("test-vault")
	v.Credentials["github"] = Credential{
		Name: "Github",
		Fields: map[string]Field{
			"username": {Label: "username", Value: "gopher", IsSecret: false},
			"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
		},
	}

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

	if c := loadedVault.Credentials["github"]; c.Name != "Github" {
		t.Errorf("Missing credential 'github' after load")
	}
}

// TestLoadMissingVault tests loading a non-existent vault
func TestLoadMissingVault(t *testing.T) {
	tmpDir := t.TempDir()
	storage := NewStorage(tmpDir)

	_, err := storage.Load("non-existent")
	if err == nil {
		t.Errorf("Expected error when loading non-existent vault, got nil")
	}
}
