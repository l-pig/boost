package cmd

import (
	"boost/internal/generator"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project components",
}

var handlerCmd = &cobra.Command{
	Use:   "handler [name]",
	Short: "Generate a new handler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		methodName := capitalize(name)

		// Default to internal/server for web projects
		// We could check if we are in a boost project, but let's assume we are at root
		outputPath := filepath.Join("internal", "server", strings.ToLower(name)+".go")

		config := generator.ComponentConfig{
			Name:         methodName,
			TemplateName: "handler",
			OutputPath:   outputPath,
		}

		fmt.Printf("Generating handler %s...\n", name)
		err := generator.GenerateComponent(config)
		if err != nil {
			fmt.Printf("Error generating handler: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Success! Generated %s\n", config.OutputPath)
		fmt.Println("Don't forget to register the handler in internal/server/server.go!")
	},
}

func init() {
	generateCmd.AddCommand(handlerCmd)
	rootCmd.AddCommand(generateCmd)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
