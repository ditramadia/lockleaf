package manager

import (
	"github.com/ditramadia/lockleaf/internal/vault"
)

type Manager struct {
	Storage vault.Storage
}

func NewManager(storage vault.Storage) *Manager {
	return &Manager{
		Storage: storage,
	}
}

// ========================================================================
// Vault handlers
// ========================================================================

func (m *Manager) CreateVault(name string) error {
	// Determine where to save the files

	// Check if vault already exists

	// Create a new vault model

	// Save the vault
}
