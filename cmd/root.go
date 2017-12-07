package cmd

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cui = rwi.New() //CUI instance
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui
	rootCmd := &cobra.Command{
		Use:   "cli-demo",
		Short: "Short comment",
		Long:  "Long comment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	rootCmd.AddCommand(newShowCmd())

	return rootCmd
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
}
