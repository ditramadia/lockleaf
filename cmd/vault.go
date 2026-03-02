package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// VaultCmd represents the base vault command
var VaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage encrypted vaults",
}

// InitCmd represents the 'vault init' command
var InitCmd = &cobra.Command{
	Use:   "init [vault_name]",
	Short: "Initialize a new vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]

		if err := globalManager.CreateVault(vaultName); err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Vault '%s' initialized.\n", vaultName)
	},
}

var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available vaults",
	Run: func(cmd *cobra.Command, args []string) {
		vaults, err := globalManager.ListVaults()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		if len(vaults) == 0 {
			fmt.Println("No vaults found. Create one with: pw vault init <name>")
			return
		}

		fmt.Println("Available Vaults:")
		for _, v := range vaults {
			fmt.Printf("  - %s\n", v)
		}
	},
}

var RenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a vault",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]
		newVaultName := args[1]

		if err := globalManager.RenameVault(vaultName, newVaultName); err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Vault '%s' renamed to '%s'.\n", vaultName, newVaultName)
	},
}

func init() {
	RootCmd.AddCommand(VaultCmd)
	VaultCmd.AddCommand(InitCmd)
	VaultCmd.AddCommand(ListCmd)
	VaultCmd.AddCommand(RenameCmd)
}
