package cmd

import (
	"boost/internal/generator"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	moduleFlag string

	templateFlag string
)

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		projectName := ""
		if len(args) > 0 {
			projectName = args[0]
		} else {
			prompt := promptui.Prompt{
				Label:   "Project Name",
				Default: "my-go-app",
				Validate: func(input string) error {
					if len(strings.TrimSpace(input)) == 0 {

						return fmt.Errorf("project name cannot be empty")

					}
					return nil
				},
			}

			projectName, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		modulePath := moduleFlag
		if modulePath == "" {
			promptMod := promptui.Prompt{
				Label:   "Module Path",
				Default: projectName,
				Validate: func(input string) error {
					if len(strings.TrimSpace(input)) == 0 {

						return fmt.Errorf("module path cannot be empty")

					}
					return nil
				},
			}
			modulePath, err = promptMod.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		templateName := templateFlag
		if templateName == "" {
			templates := []string{"basic", "web"}
			promptTmpl := promptui.Select{
				Label: "Select Template",
				Items: templates,
			}
			_, templateName, err = promptTmpl.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		config := generator.ProjectConfig{
			ProjectName: projectName,
			ModulePath:  modulePath,
			Template:    templateName,
		}

		fmt.Printf("Creating project %s...\n", projectName)
		err = generator.Generate(config)
		if err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nSuccess! Created %s at %s\n", projectName, projectName)
		fmt.Println("We suggest that you begin by typing:")
		fmt.Printf("\n  cd %s\n", projectName)
		if templateName == "web" {
			fmt.Println("  go run cmd/server/main.go")
		} else {
			fmt.Println("  go run main.go")
		}
		fmt.Println("\nHappy hacking!")
	},
}

func init() {
	createCmd.Flags().StringVarP(&moduleFlag, "module", "m", "", "Go module path")
	createCmd.Flags().StringVarP(&templateFlag, "template", "t", "", "Project template (e.g., basic)")
	rootCmd.AddCommand(createCmd)
}
