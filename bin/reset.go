package bin

import (
	"github.com/spf13/cobra"
	"loopyctl/master"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Setup the environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := master.Reset(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
