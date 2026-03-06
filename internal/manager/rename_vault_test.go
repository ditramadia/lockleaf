package manager

import (
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
)

func TestRenameVault(t *testing.T) {
	oldVaultName := "old-test-vault"
	newVaultName := "new-test-vault"

	// Setup temporary storage for the test
	tmpDir := t.TempDir()
	storage := vault.NewStorage(tmpDir)

	// Create Manager instance
	manager := NewManager(storage)

	// Inject a vault
	existingVault := vault.NewVault(oldVaultName)
	if err := manager.Storage.Save(existingVault); err != nil {
		t.Fatalf("failed to inject vault: %v", err)
	}

	// Test rename vault
	if err := manager.RenameVault(oldVaultName, newVaultName); err != nil {
		t.Errorf("failed to rename vault: %v", err)
	}

	// Load renamed vault
	v, err := manager.Storage.Load(newVaultName)
	if err != nil {
		t.Fatalf("failed to load vault: %v", err)
	}

	t.Logf("vault: %v", v)

	// Assert data
	if v.Name != newVaultName {
		t.Errorf("expected name '%s', got '%s'", newVaultName, v.Name)
	}
}
