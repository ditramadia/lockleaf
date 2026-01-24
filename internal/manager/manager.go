package manager

import (
	"errors"

	"github.com/ditramadia/lockleaf/internal/vault"
)

var (
	ErrVaultExists        = errors.New("Vault already exists")
	ErrCredentialNotFound = errors.New("Credential not found")
	ErrCredentialExists   = errors.New("Credential already exists")
	ErrFieldNotFound      = errors.New("Field not found")
)

// Manager represents the app global manager
type Manager struct {
	Storage *vault.Storage
}

// NewManager creates a new manager instance
func NewManager(storage *vault.Storage) *Manager {
	return &Manager{
		Storage: storage,
	}
}
