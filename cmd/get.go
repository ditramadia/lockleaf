package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ditramadia/lockleaf/storage"
	"github.com/spf13/cobra"
)

// @TODO: Add flag --metadata to display the created at etc.
var getCmd = &cobra.Command{
	Use:   "get [vault_name] [credential_label]",
	Short: "Get details for a specific credential",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var vaultName string = args[0]
		var searchLabel string = args[1]
		var vaultPath string = filepath.Join("vault", vaultName+".json")

		// Load vault
		vault, err := storage.LoadVault(vaultPath)
		if err != nil {
			return fmt.Errorf("Failed to load vault %s: %w", vaultName, err)
		}

		// Find the credential
		var found bool
		for _, cred := range vault.Credentials {
			if strings.EqualFold(cred.Label, searchLabel) {
				printCredential(cred)
				found = true
			}
		}

		if !found {
			return fmt.Errorf("No credential with label '%s' found", searchLabel)
		}

		return nil
	},
}

func printCredential(c storage.Credential) {
	fmt.Printf("\n--- %s ---\n", c.Label)
	fmt.Println()
	for _, field := range c.Fields {
		fmt.Printf("%s: %s\n", strings.Title(field.Key), field.Value)
	}
	fmt.Println()
	// fmt.Println("---------------------------")
}

func init() {
	RootCmd.AddCommand(getCmd)
}
