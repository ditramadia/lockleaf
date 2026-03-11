package handler

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2/list"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ditramadia/lockleaf/internal/ui"
)

func (h *Handler) InitVault(vaultName string) {
	if err := h.s.CreateVault(vaultName); err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	successMsg := fmt.Sprintf("Vault '%s' initialized.", vaultName)
	fmt.Println(ui.Success.Render(successMsg))

	os.Exit(0)
}

func (h *Handler) Connect(vaultName string) {
	h.validateVaultExists(vaultName)

	if err := h.cfg.SetActiveVault(vaultName); err != nil {
		fmt.Println(ui.Error.Render("Failed to connect."))
		os.Exit(1)
	}

	successMsg := fmt.Sprintf("Vault '%s' connected.", vaultName)
	fmt.Println(ui.Success.Render(successMsg))

	os.Exit(0)
}

func (h *Handler) ListVaults() {
	vaults, err := h.s.ListVaults()
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	if len(vaults) == 0 {
		fmt.Println(ui.Normal.MarginTop(1).Render("No vaults found."))
		fmt.Println(ui.Tips.MarginBottom(1).MarginLeft(2).Render("(Use \"leaf vault <name>\" to create a vault)"))
		os.Exit(0)
	}

	activeVault, _, err := h.cfg.GetActiveVault()
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	items := make([]any, len(vaults))
	for i, v := range vaults {
		if v == activeVault {
			items[i] = ui.Info.Bold(true).Render(v + " (active)")
		} else {
			items[i] = ui.Normal.Render(v)
		}
	}

	l := list.New(items...).
		Enumerator(func(_ list.Items, i int) string {
			return ui.BulletStyle.Render("•")
		})

	fmt.Println(ui.Normal.MarginTop(1).Render("Vaults:"))
	fmt.Println(ui.ListStyle.Render(l.String()))

	os.Exit(0)
}

func (h *Handler) RenameVault(oldName, newName string) {
	if oldName == "" {
		// Rename current vault
		activeVaultName, ok, err := h.cfg.GetActiveVault()
		if err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}
		if !ok {
			fmt.Println(ui.Error.Render("No active vault. Please connect to a vault first."))
			os.Exit(1)
		}

		if err := h.cfg.SetActiveVault(newName); err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}

		oldName = activeVaultName
	}

	if err := h.s.RenameVault(oldName, newName); err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	styledNewName := ui.Info.Bold(true).Render(newName)
	fmt.Println(ui.Success.Render(
		fmt.Sprintf("Vault '%s' renamed to '%s'.", oldName, styledNewName),
	))

	os.Exit(0)
}

func (h *Handler) DeleteVault(vaultName string, force bool) {
	h.validateVaultExists(vaultName)

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

	activeVault, ok, err := h.cfg.GetActiveVault()
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}
	if !ok {
		fmt.Println(ui.Error.Render("No active vault. Please connect to a vault first."))
		os.Exit(1)
	}

	isActive := activeVault == vaultName
	if isActive {
		if err := h.cfg.SetActiveVault(""); err != nil {
			fmt.Println(ui.Error.Render(err.Error()))
			os.Exit(1)
		}
	}

	if err := h.s.RemoveVault(vaultName); err != nil {

		// Rollback config active vault
		if isActive {
			if err := h.cfg.SetActiveVault(activeVault); err != nil {
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

// Internal helpers

func (h *Handler) validateVaultExists(vaultName string) {
	exists, err := h.s.IsVaultExist(vaultName)
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}
	if !exists {
		fmt.Println(ui.Error.Render("Vault not found."))
		fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"leaf vault\" to list vaults)"))
		os.Exit(1)
	}
}
