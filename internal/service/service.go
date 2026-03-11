package service

import (
	"errors"
	"testing"

	"github.com/ditramadia/lockleaf/internal/vault"
)

var (
	ErrVaultExists        = errors.New("Vault already exists")
	ErrVaultNotFound      = errors.New("Vault not found")
	ErrCredentialExists   = errors.New("Credential already exists")
	ErrCredentialNotFound = errors.New("Credential not found")
	ErrFieldNotFound      = errors.New("Field not found")
)

type Service struct {
	Storage *vault.Storage
}

func New(storage *vault.Storage) *Service {
	return &Service{
		Storage: storage,
	}
}

// Test helpers

// Internal helpers

func newTestStorage(t *testing.T) *vault.Storage {
	return vault.New(t.TempDir())
}

func newTestService(s *vault.Storage) *Service {
	return New(s)
}

func newTestVault(vaultName string, credentials map[string]vault.Credential) *vault.Vault {
	v := vault.NewVault(vaultName)
	if credentials != nil {
		v.Credentials = credentials
	}

	return v
}

func newTestCredentials() map[string]vault.Credential {
	c := &vault.Credential{
		Name: "Github",
		Fields: map[string]vault.Field{
			"username": {Label: "username", Value: "gopher", IsSecret: false},
			"password": {Label: "password", Value: "s3cr3t", IsSecret: true},
		},
	}

	return map[string]vault.Credential{
		c.Name: *c,
	}
}
