package manager

import (
	"errors"

	"github.com/ditramadia/lockleaf/internal/vault"
)

var (
	ErrVaultExists        = errors.New("vault already exists")
	ErrCredentialNotFound = errors.New("credential not found")
	ErrCredentialExists   = errors.New("credential already exists")
	ErrFieldNotFound      = errors.New("field not found")
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
