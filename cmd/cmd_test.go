package cmd

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestShowNormal(t *testing.T) {
	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"show", "-i", "123", "-s", "日本語", "-b"}, want: "Integer option value: 123\n String option value: 日本語\nBoolean option value: true\n"},
		{args: []string{"show", "-i", "123", "-s", "日本語"}, want: "Integer option value: 123\n String option value: 日本語\nBoolean option value: false\n"},
		{args: []string{"show", "-i", "123"}, want: "Integer option value: 123\n String option value: \nBoolean option value: false\n"},
		{args: []string{"show"}, want: "Integer option value: 0\n String option value: \nBoolean option value: false\n"},
	}

	for _, c := range testCases {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := gocli.NewUI(
			gocli.Writer(out),
			gocli.ErrorWriter(errOut),
		)
		exit := Execute(ui, c.args)
		if exit != ExitNormal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, ExitNormal)
		}
		if out.String() != c.want {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), c.want)
		}
		if errOut.String() != "" {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), "")
		}
	}
}
