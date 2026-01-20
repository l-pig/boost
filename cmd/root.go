package cmd

import (
	"boost/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boost",
	Short: "Boost is a scaffold tool for creating Go projects",
	Long:  `Boost helps you quickly set up a standardized Go project structure, similar to create-react-app.`,
	Run: func(cmd *cobra.Command, args []string) {
		program := tea.NewProgram(model.NewRoot())

		if _, err := program.Run(); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	//rootCmd.AddCommand(generateCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
