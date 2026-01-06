package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ditramadia/lockleaf/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [vault_name]",
	Short: "List all vaults or entries within a vault",
	RunE: func(cmd *cobra.Command, args []string) error {
		// List all vaults if no name is provided
		if len(args) == 0 {
			files, err := os.ReadDir("vault")
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println("No vaults found. Use 'init' to create one.")
					return nil
				}
				return err
			}

			// List vaults
			fmt.Println("Available Vaults:")
			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
					name := file.Name()[:len(file.Name())-len(".json")]
					fmt.Printf("- %s\n", name)
				}
			}
			return nil
		}

		// Load specific vault
		var vaultName string = args[0]
		var vaultPath string = filepath.Join("vault", vaultName+".json")
		vault, err := storage.LoadVault(vaultPath)
		if err != nil {
			return fmt.Errorf("Failed to load vault %s: %w", vaultName, err)
		}

		// List credentials
		count := len(vault.Credentials)
		if count == 0 {
			fmt.Printf("Vault '%s' is empty.\n", vaultName)
			return nil
		}

		fmt.Printf("Showing %d credentials in '%s':\n", count, vaultName)
		for i, credential := range vault.Credentials {
			fmt.Printf("%2d. [%-10s] %s\n", i+1, credential.Category, credential.Label)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
