package main

import (
	"os"

	"github.com/spiegel-im-spiegel/cli-demo/cmd"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func main() {
	cmd.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
