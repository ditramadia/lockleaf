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

func (s *Service) RenameCredential(vaultName, name, newName string) error {

	// Check if vault exists
	isVaultExist, err := s.Storage.IsVaultExist(vaultName)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if !isVaultExist {
		return fmt.Errorf("Vault '%s' does not exist", vaultName)
	}

	// Load the vault
	v, err := s.Storage.Load(vaultName)
	if err != nil {
		return fmt.Errorf("Error loading vault: %w", err)
	}

	// Check if credential exists
	isCredentialExist, err := s.Storage.IsCredentialExist(v, name)
	if err != nil {
		return fmt.Errorf("Error checking credential existance: %w", err)
	}
	if !isCredentialExist {
		return fmt.Errorf("Credential '%s' does not exist", name)
	}

	// Check if credential with newName already exists
	isNewCredentialExist, err := s.Storage.IsCredentialExist(v, newName)
	if err != nil {
		return fmt.Errorf("Error checking new credential existance: %w", err)
	}
	if isNewCredentialExist {
		return fmt.Errorf("Credential '%s' already exists", newName)
	}

	// Rename the credential
	credential := v.Credentials[name]
	delete(v.Credentials, name)
	credential.Name = newName
	v.Credentials[newName] = credential

	// Save the updated vault
	if err := s.Storage.Save(v); err != nil {
		return fmt.Errorf("Error saving vault: %w", err)
	}

	return nil

}
