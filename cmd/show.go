package cmd

import (
	"github.com/spf13/cobra"
)

//newShowCmd returns cobra.Command instance for show sub-command
func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Short comment for show sub-command",
		Long:  "Long comment for show sub-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := cmd.Flags().GetInt("integer")
			if err != nil {
				return err
			}
			b, err := cmd.Flags().GetBool("boolean")
			if err != nil {
				return err
			}
			s, err := cmd.Flags().GetString("string")
			if err != nil {
				return err
			}
			cui.Outputln("Integer option value:", i)
			cui.Outputln(" String option value:", s)
			cui.Outputln("Boolean option value:", b)

			return nil
		},
	}

	showCmd.Flags().IntP("integer", "i", 0, "integer option")
	showCmd.Flags().BoolP("boolean", "b", false, "boolean option")
	showCmd.Flags().StringP("string", "s", "", "string option")

	return showCmd
}
