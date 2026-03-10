package manager

import (
	"fmt"

	"github.com/ditramadia/lockleaf/internal/vault"
)

func (m *Manager) CreateCredential(vaultName, name string) error {

	// Load the vault
	v, err := m.Storage.Load(vaultName)
	if err != nil {
		return fmt.Errorf("Error loading vault: %w", err)
	}

	// Check if credential already exists
	exists, err := m.Storage.IsCredentialExist(v, name)
	if err != nil {
		return fmt.Errorf("Error checking credential existance: %w", err)
	}
	if exists {
		return fmt.Errorf("Credential '%s' already exists", name)
	}

	// Create new credential
	credential := vault.NewCredential(name)
	v.Credentials[name] = *credential

	// Save the updated vault
	if err := m.Storage.Save(v); err != nil {
		return fmt.Errorf("Error saving vault: %w", err)
	}

	return nil
}
