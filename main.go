package main

import (
	"os"

	"github.com/spiegel-im-spiegel/cli-demo/cmd"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func main() {
	cmd.Execute(
		rwi.New(
			rwi.Reader(os.Stdin),
			rwi.Writer(os.Stdout),
			rwi.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
