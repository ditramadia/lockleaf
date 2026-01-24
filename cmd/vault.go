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

func init() {
	RootCmd.AddCommand(VaultCmd)
	VaultCmd.AddCommand(InitCmd)
	VaultCmd.AddCommand(ListCmd)
}
