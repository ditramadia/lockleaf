package service

import (
	"fmt"

	"github.com/ditramadia/lockleaf/internal/vault"
)

func (s *Service) CreateVault(name string) error {
	// Check if vault already exists
	exists, err := s.Storage.IsVaultExist(name)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if exists {
		return ErrVaultExists
	}

	// Create a new vault model
	vault := vault.NewVault(name)

	// Save the vault
	if err = s.Storage.Save(vault); err != nil {
		return fmt.Errorf("Error saving the vault: %w", err)
	}

	return nil
}

func (s *Service) ListVaults() ([]string, error) {
	return s.Storage.List()
}

func (s *Service) RenameVault(name, newName string) error {
	// Check if vault exists
	exists, err := s.Storage.IsVaultExist(name)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if !exists {
		return ErrVaultNotFound
	}

	// Check if vault with newName already exists
	exists, err = s.Storage.IsVaultExist(newName)
	if err != nil {
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if exists {
		return ErrVaultExists
	}

	// Rename the vault
	if err = s.Storage.Rename(name, newName); err != nil {
		return fmt.Errorf("Error renaming the vault: %w", err)
	}

	return nil
}

func (s *Service) RemoveVault(name string) error {
	// Check if vault exists
	exists, err := s.Storage.IsVaultExist(name)
	if err != nil {
		fmt.Println("Confirmation failed.")
		fmt.Println("Aborted.")
		return fmt.Errorf("Error checking vault existance: %w", err)
	}
	if !exists {
		return ErrVaultNotFound
	}

	// Remove the vault
	if err = s.Storage.Remove(name); err != nil {
		return fmt.Errorf("Error removing the vault: %w", err)
	}

	return nil
}

func (s *Service) IsVaultExist(name string) (bool, error) {
	// Check if vault exists
	exists, err := s.Storage.IsVaultExist(name)
	if err != nil {
		return false, fmt.Errorf("Error checking vault existance: %w", err)
	}

	return exists, nil
}
