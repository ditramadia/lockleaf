package service

import (
	"fmt"

	"github.com/ditramadia/lockleaf/internal/vault"
)

func (s *Service) CreateCredential(vaultName, name string) error {

	// Load the vault
	v, err := s.Storage.Load(vaultName)
	if err != nil {
		return fmt.Errorf("Error loading vault: %w", err)
	}

	// Check if credential already exists
	exists, err := s.Storage.IsCredentialExist(v, name)
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
	if err := s.Storage.Save(v); err != nil {
		return fmt.Errorf("Error saving vault: %w", err)
	}

	return nil
}

func (s *Service) ListCredentials(vaultName string) ([]string, error) {

	// Load the vault
	v, err := s.Storage.Load(vaultName)
	if err != nil {
		return nil, fmt.Errorf("Error loading vault: %w", err)
	}

	var credentials []string
	for _, cred := range v.Credentials {
		credentials = append(credentials, cred.Name)
	}
	return credentials, nil
}
