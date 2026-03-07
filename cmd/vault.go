package cmd

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2/list"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ditramadia/lockleaf/internal/config"
	"github.com/ditramadia/lockleaf/internal/ui"
	"github.com/spf13/cobra"
)

var connect string
var delete string
var force bool

var vaultCmd = &cobra.Command{
	Use:     "vault",
	Aliases: []string{"v"},
	Short:   "Manage encrypted vaults",
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to a vault
		if connect != "" {
			vaultName := connect

			exists, err := globalManager.IsVaultExist(vaultName)
			if err != nil {
				fmt.Println(ui.Error.Render(err.Error()))
				os.Exit(1)
			}
			if !exists {
				fmt.Println(ui.Error.Render("Vault not found."))
				fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"leaf vault\" to list vaults)"))
				os.Exit(1)
			}

			if err := config.SetActiveVault(vaultName); err != nil {
				fmt.Println(ui.Error.Render("Failed to connect."))
				os.Exit(1)
			}

			successMsg := fmt.Sprintf("Vault '%s' connected.", vaultName)
			fmt.Println(ui.Success.Render(successMsg))
			return
		}

		// Delete a vault
		if delete != "" {
			vaultName := delete

			exists, err := globalManager.IsVaultExist(vaultName)
			if err != nil {
				fmt.Println(ui.Error.Render(err.Error()))
				os.Exit(1)
			}
			if !exists {
				fmt.Println(ui.Error.Render("Vault not found."))
				fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"leaf vault\" to list vaults)"))
				os.Exit(1)
			}

			// Prompt user for confirmation
			if !force {

				confirm := ""
				match := "vault/" + vaultName

				promptMsg := fmt.Sprintf("Type '%s' to confirm:", match)
				prompt := &survey.Input{
					Message: fmt.Sprint(ui.Normal.Render(promptMsg)),
				}
				survey.AskOne(prompt, &confirm,
					survey.WithStdio(os.Stdin, os.Stderr, os.Stderr),
					survey.WithIcons(func(icons *survey.IconSet) {
						icons.Question.Format = "reset"
						icons.Question.Text = "?"
					}),
				)

				if confirm != match {
					fmt.Println(ui.Error.Render("Confirmation failed."))
					os.Exit(0)
				}
			}

			activeVault := config.GetActiveVault()
			isDeleteCurrentVault := activeVault == delete
			if isDeleteCurrentVault {
				if err := config.SetActiveVault(""); err != nil {
					fmt.Println(ui.Error.Render(err.Error()))
					os.Exit(1)
				}
			}

			if err := globalManager.RemoveVault(vaultName); err != nil {

				// Rollback config active vault
				if isDeleteCurrentVault {
					if err := config.SetActiveVault(activeVault); err != nil {
						fmt.Println(ui.Error.Render(err.Error()))
						os.Exit(1)
					}
				}

				fmt.Println(ui.Error.Render(err.Error()))
				os.Exit(1)
			}

			successMsg := fmt.Sprintf("Vault '%s' removed.", vaultName)
			fmt.Println(ui.Success.Render(successMsg))

			os.Exit(0)
		}

		// Initialize a new vault
		if len(args) > 0 {
			vaultName := args[0]

			if err := globalManager.CreateVault(vaultName); err != nil {
				fmt.Println(ui.Error.Render(err.Error()))
				os.Exit(1)
			}

			successMsg := fmt.Sprintf("Vault '%s' initialized.", vaultName)
			fmt.Println(ui.Success.Render(successMsg))

			os.Exit(0)
		}

		// List all available vaults
		{
			vaults, err := globalManager.ListVaults()
			if err != nil {
				fmt.Println(ui.Error.Render(err.Error()))
				os.Exit(1)
			}

			if len(vaults) == 0 {
				fmt.Println(ui.Normal.Render("No vaults found."))
				fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"leaf vault init <name>\" to create a vault)"))
				os.Exit(1)
			}

			activeVault := config.GetActiveVault()
			items := make([]any, len(vaults))
			for i, v := range vaults {
				// 2. Check if this is the active one
				if v == activeVault {
					// Style ONLY this item
					items[i] = ui.Info.Bold(true).Render(v + " (active)")
				} else {
					// Keep the rest normal
					items[i] = ui.Normal.Render(v)
				}
			}

			l := list.New(items...).
				Enumerator(func(_ list.Items, i int) string {
					return ui.BulletStyle.Render("•")
				})

			fmt.Println(ui.Normal.MarginTop(1).Render("Available Vaults:"))
			fmt.Println(ui.ListStyle.Render(l.String()))

			os.Exit(0)
		}
	},
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a vault",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]
		newVaultName := args[1]

		if err := globalManager.RenameVault(vaultName, newVaultName); err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}

		styledNewName := ui.Info.Bold(true).Render(newVaultName)
		fmt.Println(ui.Success.Render(
			fmt.Sprintf("Vault '%s' renamed to '%s'.", vaultName, styledNewName),
		))
	},
}

func init() {
	vaultCmd.Flags().StringVarP(&connect, "connect", "c", "", "Connect to a specific vault")
	vaultCmd.Flags().StringVarP(&delete, "delete", "d", "", "Delete a vault")
	vaultCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation prompt")

	rootCmd.AddCommand(vaultCmd)
	vaultCmd.AddCommand(renameCmd)
}
