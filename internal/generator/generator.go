package generator

import (
	"boost/internal/templates"
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ComponentConfig struct {
	Name         string
	TemplateName string
	OutputPath   string
}

func GenerateComponent(config ComponentConfig) error {
	templatePath := "components/" + config.TemplateName + ".go.tmpl"

	// Read template
	content, err := templates.Templates.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Parse template
	tmpl, err := template.New(templatePath).Parse(string(content))
	if err != nil {
		return err
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return err
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(config.OutputPath), 0755); err != nil {
		return err
	}

	// Write file
	return os.WriteFile(config.OutputPath, buf.Bytes(), 0644)
}

type ProjectConfig struct {
	ProjectName string
	ModulePath  string
	Template    string
}

func Generate(config ProjectConfig) error {
	targetDir := config.ProjectName
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	templateDir := config.Template
	return fs.WalkDir(templates.Templates, templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Read template
		content, err := templates.Templates.ReadFile(path)
		if err != nil {
			return err
		}

		// Parse template
		tmpl, err := template.New(path).Parse(string(content))
		if err != nil {
			return err
		}

		// Prepare output path
		relPath, _ := filepath.Rel(templateDir, path)
		outputPath := filepath.Join(targetDir, strings.TrimSuffix(relPath, ".tmpl"))

		// Create directories for output
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		// Execute template
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, config); err != nil {
			return err
		}

		// Write file
		return os.WriteFile(outputPath, buf.Bytes(), 0644)
	})
}
