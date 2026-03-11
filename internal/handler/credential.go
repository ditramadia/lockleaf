package handler

import (
	"fmt"
	"os"

	"github.com/ditramadia/lockleaf/internal/ui"
)

func (h *Handler) AddCredential(credentialName string) {
	// Get current active vault
	activeVault, ok, err := h.cfg.GetActiveVault()
	if err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}
	if !ok {
		fmt.Println(ui.Error.Render("No active vault. Please connect to a vault first."))
		os.Exit(1)
	}

	// Create credential in the active vault
	if err := h.m.CreateCredential(activeVault, credentialName); err != nil {
		fmt.Println(ui.Error.Render(err.Error()))
		os.Exit(1)
	}

	successMsg := fmt.Sprintf("Credential '%s' added.", credentialName)
	fmt.Println(ui.Success.Render(successMsg))

	os.Exit(0)
}
