package cmd

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2/list"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ditramadia/lockleaf/internal/ui"
	"github.com/spf13/cobra"
)

var forceDelete bool

// VaultCmd represents the base vault command
var vaultCmd = &cobra.Command{
	Use:     "vault",
	Aliases: []string{"v"},
	Short:   "Manage encrypted vaults",
}

// InitCmd represents the 'vault init' command
var initCmd = &cobra.Command{
	Use:   "init [vault_name]",
	Short: "Initialize a new vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]

		if err := globalManager.CreateVault(vaultName); err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}

		successMsg := fmt.Sprintf("Vault '%s' initialized.", vaultName)
		fmt.Println(ui.Success.Render(successMsg))
	},
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available vaults",
	Run: func(cmd *cobra.Command, args []string) {
		vaults, err := globalManager.ListVaults()
		if err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}

		if len(vaults) == 0 {
			fmt.Println(ui.Normal.Render("No vaults found."))
			fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"pw vault init <name>\" to create a vault)"))
			os.Exit(1)
		}

		items := make([]any, len(vaults))
		for i, v := range vaults {
			items[i] = v
		}

		l := list.New(items...).
			Enumerator(func(_ list.Items, i int) string {
				return ui.BulletStyle.Render("•")
			}).
			ItemStyle(ui.Normal)

		fmt.Println(ui.Normal.MarginTop(1).Render("Available Vaults:"))
		fmt.Println(ui.ListStyle.Render(l.String()))
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

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]

		exists, err := globalManager.IsVaultExist(vaultName)
		if err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}
		if !exists {
			fmt.Println(ui.Error.Render("Vault not found."))
			fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"pw vault ls\" to list vaults)"))
			os.Exit(1)
		}

		// Prompt user for confirmation
		if !forceDelete {

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

		if err := globalManager.RemoveVault(vaultName); err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}

		successMsg := fmt.Sprintf("Vault '%s' removed.", vaultName)
		fmt.Println(ui.Success.Render(successMsg))
	},
}

func init() {
	removeCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation prompt")

	rootCmd.AddCommand(vaultCmd)
	vaultCmd.AddCommand(initCmd)
	vaultCmd.AddCommand(listCmd)
	vaultCmd.AddCommand(renameCmd)
	vaultCmd.AddCommand(removeCmd)
}
