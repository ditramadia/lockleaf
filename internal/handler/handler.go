package handler

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ditramadia/lockleaf/internal/config"
	"github.com/ditramadia/lockleaf/internal/service"
	"github.com/ditramadia/lockleaf/internal/ui"
)

type Handler struct {
	cfg *config.Config
	s   *service.Service
}

func New(config *config.Config, service *service.Service) *Handler {
	return &Handler{
		cfg: config,
		s:   service,
	}
}

func (h *Handler) askConfirmation(match string) {
	confirm := ""

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
