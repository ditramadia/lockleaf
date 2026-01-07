package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ditramadia/lockleaf/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [vault_name]",
	Short: "Add a new credential into a vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultName := args[0]
		vaultPath := filepath.Join("vault", vaultName+".json")

		// Load the vault
		vault, err := storage.LoadVault(vaultPath)
		if err != nil {
			return fmt.Errorf("Failed to load vault %s: %w", vaultName, err)
		}

		// Collect header info
		var label, category string
		fmt.Print("Credential Label (e.g. Facebook): ")
		fmt.Scanln(&label)
		fmt.Print("Category (e.g. Social Media): ")
		fmt.Scanln(&category)

		// Create new credential
		newCredential := storage.Credential{
			ID:       fmt.Sprintf("%d", time.Now().Unix()),
			Label:    label,
			Category: category,
			Fields:   []storage.Field{},
		}

		// Field collection loop
		for {
			fmt.Println("\nAdd a field:")
			fmt.Println("1. Username")
			fmt.Println("2. Email")
			fmt.Println("3. Password")
			fmt.Println("4. Phone")
			fmt.Println("5. PIN")
			fmt.Println("0. Finish")
			fmt.Print("Select an option: ")

			var choice int
			fmt.Scanln(&choice)

			if choice == 0 {
				break
			}

			var fType storage.FieldType
			var key string

			switch choice {
			case 1:
				fType, key = storage.FieldUsername, "username"
			case 2:
				fType, key = storage.FieldEmail, "email"
			case 3:
				fType, key = storage.FieldPassword, "password"
			case 4:
				fType, key = storage.FieldPhone, "phone"
			case 5:
				fType, key = storage.FieldPIN, "pin"
			default:
				fmt.Println("Invalid choice, try again.")
				continue
			}

			fmt.Printf("Enter %s value: ", key)
			var value string
			fmt.Scanln(&value)

			newCredential.Fields = append(newCredential.Fields, storage.Field{
				Key:   key,
				Type:  fType,
				Value: value,
			})
		}

		// Update and save
		vault.Credentials = append(vault.Credentials, newCredential)
		err = storage.SaveVault(vault, vaultPath)
		if err != nil {
			return fmt.Errorf("Failed to save vault %s: %w", vaultName, err)
		}

		fmt.Printf("\nSuccessfully added '%s' to '%s'!\n", label, vaultName)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
