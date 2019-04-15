package check

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	cmdutil "github.com/morozovcookie/goproject-cli/pkg/cmd/util"
)

type Options struct {
	GitPath string
}

func NewCheckOptions() *Options {
	return &Options{}
}

func NewCheckCommand() (command *cobra.Command) {
	opts := NewCheckOptions()

	command = &cobra.Command{
		Use:   "check",
		Short: "Checking dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(opts.Run())
		},
	}

	return command
}

func (v *Options) Run() (err error) {
	if v.GitPath, err = v.findExecutable("git"); err != nil {
		return err
	}

	return nil
}

func (v Options) findExecutable(name string) (path string, err error) {
	_, _ = fmt.Fprintf(os.Stdout, "trying to find \"%s\". . .\n", name)

	if path, err = exec.LookPath(name); err != nil {
		return "", err
	}

	_, _ = fmt.Fprintf(os.Stdout, "\"%s\" was found in %s\n", name, path)

	return path, nil
}
