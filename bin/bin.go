package bin

import (
	"fmt"
	"github.com/spf13/cobra"
	"loopyctl/pkg/logger"
	"os"
)

var (
	cfgFile string
	Info    bool
)

var rootCmd = &cobra.Command{
	Use:   "loopyctl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.loopyctl/config.yaml)")
	rootCmd.PersistentFlags().BoolVar(&Info, "info", false, "set logger level to Info (default is Debug)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}

	logFile := fmt.Sprintf("%s/.loopyctl/loopyctl.log", home)
	if !FileExist(home + "/.loopyctl") {
		err := os.MkdirAll(home+"/.loopyctl", os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create config directory: %v\n", err)
			fmt.Println("Please create it manually with the command: mkdir -p /root/.loopyctl && touch /root/.loopyctl/config.yaml")
			os.Exit(1)
		}
	}

	// Set the logger configuration based on the Info flag.
	if Info {
		logger.Cfg(5, logFile) // Assuming 5 is the log level for Info.
	} else {
		logger.Cfg(6, logFile) // Assuming 6 is the log level for Debug.
	}
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
