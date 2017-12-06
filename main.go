package main

import (
	"os"

	"github.com/spiegel-im-spiegel/cli-demo/cmd"
	"github.com/spiegel-im-spiegel/gocli"
)

func main() {
	os.Exit(cmd.Execute(
		gocli.NewUI(
			gocli.Reader(os.Stdin),
			gocli.Writer(os.Stdout),
			gocli.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Int())
}
