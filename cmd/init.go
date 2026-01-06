package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/ditramadia/lockleaf/storage"
	"github.com/spf13/cobra"
)

// Command to initialize a new vault by providing the name of the vault. The vault file will be saved by default in 'vault/'
// Example: pw init [vault_name]
var initCmd = &cobra.Command{
	Use:   "init [vault_name]",
	Short: "Initialize a new vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var vaultName string = args[0]
		var vaultPath string = filepath.Join("vault", vaultName+".json")

		// Check if vault already exists
		if storage.VaultExists(vaultPath) {
			return fmt.Errorf("A vault named '%s' already exists at %s", vaultName, vaultPath)
		}

		// Initialize the vault's struct
		vault := storage.Vault{
			Name:        vaultName,
			Credentials: []storage.Credential{},
		}

		// Save vault as .json (for now)
		err := storage.SaveVault(vault, vaultPath)
		if err != nil {
			return fmt.Errorf("Failed to create vault: %w", err)
		}

		fmt.Printf("Vault '%s' initialized at %s\n", vaultName, vaultPath)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
