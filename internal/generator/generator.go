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

type ProjectConfig struct {
	ProjectName  string
	Module       string
	Dependencies []string
}

func Generate(config ProjectConfig) error {
	targetDir := config.ProjectName
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	templateDir := "web"
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
