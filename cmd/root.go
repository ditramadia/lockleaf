package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ditramadia/lockleaf/internal/config"
	"github.com/ditramadia/lockleaf/internal/handler"
	"github.com/ditramadia/lockleaf/internal/manager"
	"github.com/ditramadia/lockleaf/internal/vault"
	"github.com/spf13/cobra"
)

var customPath string
var h *handler.Handler

var rootCmd = &cobra.Command{
	Use:        "leaf",
	Short:      "Lockleaf: A secure CLI password manager",
	SuggestFor: []string{"pw"},
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
	// Determine storage directory
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

	// Initialize config
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("error initializing config: %v", err)
	}

	// Setup storage
	storage := vault.New(dataDir)

	// Setup manager
	m := manager.New(storage)

	// Setup handlers
	h = handler.New(cfg, m)

	return nil
}
