package cmd

import (
	"github.com/spf13/cobra"
)

var connect string
var delete string
var modify string
var force bool

var vaultCmd = &cobra.Command{
	Use:     "vault",
	Aliases: []string{"v"},
	Short:   "Manage encrypted vaults",
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to a vault
		if connect != "" {
			vaultName := connect
			h.Connect(vaultName)
		}

		// Delete a vault
		if delete != "" {
			vaultName := delete
			h.DeleteVault(vaultName, force)
		}

		// Rename a vault
		if modify != "" {
			newVaultName := modify
			var vaultName string

			// Rename another vault
			if len(args) == 1 {
				vaultName = args[0]
			}

			h.RenameVault(vaultName, newVaultName)
		}

		// Initialize a new vault
		if len(args) == 1 {
			vaultName := args[0]
			h.InitVault(vaultName)
		}

		// List all available vaults
		h.ListVaults()
	},
}

func init() {
	vaultCmd.Flags().StringVarP(&connect, "connect", "c", "", "Connect to a specific vault")
	vaultCmd.Flags().StringVarP(&delete, "delete", "d", "", "Delete a vault")
	vaultCmd.Flags().StringVarP(&modify, "modify", "m", "", "Rename a vault")
	vaultCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation prompt")

	rootCmd.AddCommand(vaultCmd)
}
