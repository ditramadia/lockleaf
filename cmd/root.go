package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pw",
	Short: "A secure CLI password manager",
}

func Execute() error {
	return RootCmd.Execute()
}
