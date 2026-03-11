package cmd

import "github.com/spf13/cobra"

var modifyCredential string
var deleteCredential string
var forceCredential bool

var credCmd = &cobra.Command{
	Use:     "cred",
	Aliases: []string{"c"},
	Short:   "Manage credentials within the active vault",
	Run: func(cmd *cobra.Command, args []string) {

		// Delete a credential
		if deleteCredential != "" {
			credentialName := deleteCredential
			h.DeleteCredential(credentialName, forceCredential)
			return
		}

		// Rename a credential
		if modifyCredential != "" {
			var oldCredentialName string

			if len(args) == 1 {
				oldCredentialName = args[0]
			}

			newCredentialName := modifyCredential
			h.RenameCredential(oldCredentialName, newCredentialName)
			return
		}

		// Add a new credential
		if len(args) == 1 {
			credentialName := args[0]
			h.AddCredential(credentialName)
			return
		}

		// List all credentials in the active vault
		h.ListCredentials()
	},
}

func init() {
	credCmd.Flags().StringVarP(&modifyCredential, "modify", "m", "", "Rename a credential")
	credCmd.Flags().StringVarP(&deleteCredential, "delete", "d", "", "Delete a credential")
	credCmd.Flags().BoolVarP(&forceCredential, "force", "f", false, "Skip confirmation prompt")

	rootCmd.AddCommand(credCmd)
}
