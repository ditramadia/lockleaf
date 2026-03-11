package cmd

import "github.com/spf13/cobra"

var credCmd = &cobra.Command{
	Use:     "cred",
	Aliases: []string{"c"},
	Short:   "Manage credentials within the active vault",
	Run: func(cmd *cobra.Command, args []string) {

		// Add a new credential
		if len(args) == 1 {
			credentialName := args[0]
			h.AddCredential(credentialName)
		}

		// List all credentials in the active vault
		h.ListCredentials()

	},
}

func init() {
	rootCmd.AddCommand(credCmd)
}
