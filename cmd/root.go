package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ditramadia/lockleaf/internal/manager"
	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/spf13/cobra"
)

var globalManager *manager.Manager

var RootCmd = &cobra.Command{
	Use:   "pw",
	Short: "Lockleaf: A secure CLI password manager",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := Setup()
		if err != nil {
			return fmt.Errorf("Error setting up cobra commands: %v\n", err)
		}
		return nil
	},
}

func Execute() error {
	return RootCmd.Execute()
}

func Setup() error {
	// Config
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("Error finding storage directory: %v\n", err)
	}
	dataDir := filepath.Join(baseDir, "lockleaf")
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return fmt.Errorf("Error creating storage directory: %v\n", err)
	}

	// Setup storage
	storage := vault.NewStorage(dataDir)

	// Setup manager
	globalManager.Storage = storage

	return nil
}
