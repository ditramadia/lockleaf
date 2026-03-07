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
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if exists {
		return ErrVaultExists
	}

	// Create a new vault model
	vault := vault.NewVault(name)

	// Save the vault
	if err = m.Storage.Save(vault); err != nil {
		return fmt.Errorf("Error saving the vault: %w", err)
	}

	return nil
}

// ListVaults returns all created vaults
func (m *Manager) ListVaults() ([]string, error) {
	return m.Storage.List()
}

// RenameVault renames a vault
func (m *Manager) RenameVault(name, newName string) error {
	// Check if vault exists
	exists, err := m.Storage.IsVaultExists(name)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if !exists {
		return ErrVaultNotFound
	}

	// Check if vault with newName already exists
	exists, err = m.Storage.IsVaultExists(newName)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if exists {
		return ErrVaultExists
	}

	// Rename the vault
	if err = m.Storage.Rename(name, newName); err != nil {
		return fmt.Errorf("Error renaming the vault: %w", err)
	}

	return nil
}
