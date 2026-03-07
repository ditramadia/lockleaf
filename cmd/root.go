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
var customPath string

var rootCmd = &cobra.Command{
	Use:   "pw",
	Short: "Lockleaf: A secure CLI password manager",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := Setup()
		if err != nil {
			return fmt.Errorf("error setting up cobra commands: %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&customPath, "path", "p", "", "custom path for vault files")
}

func Execute() error {
	return rootCmd.Execute()
}

func Setup() error {
	// Config
	var dataDir string
	if customPath != "" {
		dataDir = customPath
	} else {
		// Default logic using os.UserConfigDir()
		baseDir, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("error finding storage directory: %v", err)
		}
		dataDir = filepath.Join(baseDir, "lockleaf", "vaults")
	}

	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return fmt.Errorf("error creating storage directory: %v", err)
	}

	// Setup storage
	storage := vault.NewStorage(dataDir)

	// Setup manager
	globalManager = manager.NewManager(storage)

	return nil
}
