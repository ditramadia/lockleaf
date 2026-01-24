package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// vaultCmd represents the base vault command
var VaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage encrypted vaults",
}

// initCmd represents the 'vault init' command
var InitCmd = &cobra.Command{
	Use:   "init [vault_name]",
	Short: "Initialize a new vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]

		err := globalManager.CreateVault(vaultName)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Vault '%s' initialized.\n", vaultName)
	},
}

func init() {
	RootCmd.AddCommand(VaultCmd)
	VaultCmd.AddCommand(InitCmd)
}
