package check

import (
	"github.com/spf13/cobra"
)

func NewCheckCommand() (command *cobra.Command) {
	command = &cobra.Command{
		Use:   "check",
		Short: "Checking dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return command
}
