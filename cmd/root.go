package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boost",
	Short: "Boost is a scaffold tool for creating Go projects",
	Long:  `Boost helps you quickly set up a standardized Go project structure, similar to create-react-app.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Root flags if any
}
