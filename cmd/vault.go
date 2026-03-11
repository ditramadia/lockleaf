package cmd

import (
	"github.com/spf13/cobra"
)

var connect string
var modifyVault string
var deleteVault string
var forceVault bool

var vaultCmd = &cobra.Command{
	Use:     "vault",
	Aliases: []string{"v"},
	Short:   "Manage encrypted vaults",
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to a vault
		if connect != "" {
			vaultName := connect
			h.Connect(vaultName)
			return
		}

		// Delete a vault
		if deleteVault != "" {
			vaultName := deleteVault
			h.DeleteVault(vaultName, forceVault)
			return
		}

		// Rename a vault
		if modifyVault != "" {
			newVaultName := modifyVault
			var vaultName string

			// Rename another vault
			if len(args) == 1 {
				vaultName = args[0]
			}

			h.RenameVault(vaultName, newVaultName)
			return
		}

		// Initialize a new vault
		if len(args) == 1 {
			vaultName := args[0]
			h.InitVault(vaultName)
			return
		}

		// List all available vaults
		h.ListVaults()
	},
}

func init() {
	vaultCmd.Flags().StringVarP(&connect, "connect", "c", "", "Connect to a specific vault")
	vaultCmd.Flags().StringVarP(&deleteVault, "delete", "d", "", "Delete a vault")
	vaultCmd.Flags().StringVarP(&modifyVault, "modify", "m", "", "Rename a vault")
	vaultCmd.Flags().BoolVarP(&forceVault, "force", "f", false, "Skip confirmation prompt")

	rootCmd.AddCommand(vaultCmd)
}
