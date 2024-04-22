package bin

import (
	"github.com/spf13/cobra"
	"loopyctl/master"
	"os"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Setup the environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if master.Single {
			// Run in single mode
			return master.Reset()
		}
		err := cmd.Help()
		if err != nil {
			return err
		}
		os.Exit(1)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
	setupCmd.Flags().BoolVar(&master.Single, "single", false, "Run in single mode")
}
