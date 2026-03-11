package cmd

import "github.com/spf13/cobra"

var modifyCredential string

var credCmd = &cobra.Command{
	Use:     "cred",
	Aliases: []string{"c"},
	Short:   "Manage credentials within the active vault",
	Run: func(cmd *cobra.Command, args []string) {

		// Rename a credential
		if modifyCredential != "" {
			var oldCredentialName string

			if len(args) == 1 {
				oldCredentialName = args[0]
			}

			newCredentialName := modifyCredential
			h.RenameCredential(oldCredentialName, newCredentialName)
		}

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
	credCmd.Flags().StringVarP(&modifyCredential, "modify", "m", "", "Rename a credential")

	rootCmd.AddCommand(credCmd)
}
