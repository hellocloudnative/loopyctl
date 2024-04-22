package bin

import (
	"github.com/spf13/cobra"
	"loopyctl/master"
	"os"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if master.Single {
			// Run in single mode
			return master.RunSetup()
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
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().BoolVar(&master.Single, "single", false, "Run in single mode")
}
