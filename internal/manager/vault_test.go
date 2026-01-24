package manager

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
)

// TestCreateVault tests the creation of a vault
func TestCreateVault(t *testing.T) {
	vaultName := "test-vault"

	// Setup temporary storage for the test
	tmpDir := t.TempDir()
	storage := vault.NewStorage(tmpDir)

	// Create Manager instance
	manager := NewManager(storage)

	// Test create vault
	err := manager.CreateVault(vaultName)
	if err != nil {
		t.Fatalf("Failed to create vault: %v", err)
	}

	// Load newly created vault
	v, err := manager.Storage.Load(vaultName)
	if err != nil {
		t.Errorf("Failed to load vault: %v", err)
	}

	// Assert data
	if v.Name != vaultName {
		t.Errorf("Expected name test-vault, got %s", v.Name)
	}
}

// TestCreateExistingVault tests the creation of an already existing vault
func TestCreateExistingVault(t *testing.T) {
	vaultName := "test-vault"

	// Setup temporary storage for the test
	tmpDir := t.TempDir()
	storage := vault.NewStorage(tmpDir)

	// Create Manager instance
	manager := NewManager(storage)

	// Inject a vault
	existingVault := vault.NewVault(vaultName)
	manager.Storage.Save(existingVault)

	// Test create vault
	err := manager.CreateVault(vaultName)
	if err == nil {
		t.Fatalf("Expected error when creating an existing vault, got nil")
	}
}
