package handler

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2/list"
	"github.com/ditramadia/lockleaf/internal/ui"
)

func (h *Handler) AddCredential(credentialName string) {

	activeVault := h.getActiveVault()

	if err := h.s.CreateCredential(activeVault, credentialName); err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	successMsg := fmt.Sprintf("Credential '%s' added.", credentialName)
	fmt.Println(ui.Success.Render(successMsg))

	os.Exit(0)
}

func (h *Handler) ListCredentials() {

	activeVault := h.getActiveVault()

	credentials, err := h.s.ListCredentials(activeVault)
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	fmt.Println(ui.Info.Bold(true).MarginTop(1).Render("Vault: " + activeVault))

	if len(credentials) == 0 {
		fmt.Println(ui.Normal.Render("No credentials found."))
		fmt.Println(ui.Tips.MarginBottom(1).MarginLeft(2).Render("(Use \"leaf cred <name>\" to add a credential)"))
		os.Exit(0)
	}

	items := make([]any, len(credentials))
	for i, v := range credentials {
		items[i] = ui.Normal.Render(v)
	}

	l := list.New(items...).
		Enumerator(func(_ list.Items, i int) string {
			return ui.BulletStyle.Render(fmt.Sprintf("%d.", i+1))
		})

	fmt.Println(ui.Normal.Render("Credentials:"))
	fmt.Println(ui.ListStyle.Render(l.String()))

	os.Exit(0)
}

func (h *Handler) RenameCredential(oldName, newName string) {
	if oldName == "" {
		fmt.Println(ui.Error.Render("Please specify the credential to rename."))
		fmt.Println(ui.Tips.MarginLeft(2).Render("(Use \"leaf cred -m <new-name> <old-name>\" to rename a credential)"))
		os.Exit(1)
	}

	activeVault := h.getActiveVault()

	if err := h.s.RenameCredential(activeVault, oldName, newName); err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	styledNewName := ui.Info.Bold(true).Render(newName)
	fmt.Println(ui.Success.Render(
		fmt.Sprintf("Credential '%s' renamed to '%s'.", oldName, styledNewName),
	))

	os.Exit(0)
}

// Internal helpers
func (h *Handler) getActiveVault() string {
	activeVault, ok, err := h.cfg.GetActiveVault()
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}
	if !ok {
		fmt.Println(ui.Error.Render("No active vault. Please connect to a vault first."))
		os.Exit(1)
	}

	return activeVault
}
