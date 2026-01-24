package manager

import (
	"fmt"

	"github.com/ditramadia/lockleaf/internal/vault"
)

// CreateVault creates a new vault
func (m *Manager) CreateVault(name string) error {
	// Check if vault already exists
	exists, err := m.Storage.IsVaultExists(name)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %v", err)
	}
	if exists {
		return ErrVaultExists
	}

	// Create a new vault model
	vault := vault.NewVault(name)

	// Save the vault
	err = m.Storage.Save(vault)
	if err != nil {
		return fmt.Errorf("Error saving the vault: %v", err)
	}

	return nil
}

// ListVaults returns all created vaults
func (m *Manager) ListVaults() ([]string, error) {
	return m.Storage.List()
}
