package cmd

import "github.com/spf13/cobra"

// vaultCmd represents the base vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage encrypted vaults",
}

// initCmd represents the 'vault init' command
var initCmd = &cobra.Command{
	Use:   "init [vault_name]",
	Short: "Initialize a new vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]
		globalManager.CreateVault(vaultName)
	},
}
