package cmd

import (
	"github.com/spf13/cobra"

	checkCmd "github.com/morozovcookie/goproject-cli/pkg/cmd/check"
	initCmd "github.com/morozovcookie/goproject-cli/pkg/cmd/init"
)

func NewDefaultCommand() (command *cobra.Command) {
	command = &cobra.Command{
		Use: "goproject-cli",
		Short: "goproject-cli as a tool for generating project skeleton",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	command.AddCommand(checkCmd.NewCheckCommand())
	command.AddCommand(initCmd.NewInitCommand())

	return command
}
