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
	if err := manager.CreateVault(vaultName); err != nil {
		t.Fatalf("failed to create vault: %v", err)
	}

	// Load newly created vault
	v, err := manager.Storage.Load(vaultName)
	if err != nil {
		t.Fatalf("failed to load vault: %v", err)
	}

	// Assert data
	if v.Name != vaultName {
		t.Errorf("expected name '%s', got '%s'", vaultName, v.Name)
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
	if err := manager.Storage.Save(existingVault); err != nil {
		t.Fatalf("failed to inject vault: %v", err)
	}

	// Test create vault
	if err := manager.CreateVault(vaultName); err == nil {
		t.Errorf("expected error when creating an existing vault, got nil")
	}
}
