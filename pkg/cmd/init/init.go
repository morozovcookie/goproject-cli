package init

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"regexp"

	cmdutil "github.com/morozovcookie/goproject-cli/pkg/cmd/util"
)

type Options struct {
	moduleName string
}

func NewInitOptions() *Options {
	return &Options{}
}

func NewInitCommand() (command *cobra.Command) {
	opts := NewInitOptions()

	command = &cobra.Command{
		Use:   "init [MODULE]",
		Short: "Init project",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(opts.Complete(cmd, args))
			cmdutil.CheckErr(opts.Validate())
			cmdutil.CheckErr(opts.Run())
		},
	}

	return command
}

func (v *Options) Complete(cmd *cobra.Command, args []string) (err error) {
	_, _ = fmt.Fprintln(os.Stdout, "reading command parameters. . .")

	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	v.moduleName = args[0]

	return nil
}

func (v *Options) Validate() (err error) {
	_, _ = fmt.Fprintln(os.Stdout, "validating command parameters. . .")

	var (
		moduleNameRegExp *regexp.Regexp
	)

	moduleNameRegExp, err = regexp.Compile(`^((?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9])/(\S+)/(\S)$`)
	if err != nil {
		return err
	}

	if !moduleNameRegExp.MatchString(v.moduleName) {
		return errors.Errorf(`invalid module name "%s"`, v.moduleName)
	}

	return nil
}

func (v *Options) Run() (err error) {
	_, _ = fmt.Fprintln(os.Stdout, `running "init" command. . .`)

	return nil
}
