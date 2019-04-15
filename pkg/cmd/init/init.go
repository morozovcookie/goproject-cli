package init

import (
	"github.com/spf13/cobra"
)

func NewInitCommand() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "init",
		Short: "Init project",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return command
}
