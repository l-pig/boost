package generator

import (
	"boost/internal/templates"
	"bytes"
	"fmt"
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

	isGin    bool
	isFiber  bool
	hasViper bool
	hasGorm  bool
	hasBob   bool
}

func (p ProjectConfig) verify() {
	for _, item := range p.Dependencies {
		if strings.Contains(item, "github.com/gin-gonic/gin") {
			p.isGin = true
		} else if strings.Contains(item, "github.com/gofiber/fiber") {
			p.isFiber = true
		} else if strings.Contains(item, "github.com/spf13/viper") {
			p.hasViper = true
		} else if strings.Contains(item, "ggorm.io/gorm") {
			p.hasGorm = true
		} else if strings.Contains(item, "github.com/stephenafamo/bob") {
			p.hasBob = true
		}
	}
}

func Generate(config ProjectConfig) error {

	config.verify()

	targetDir := config.ProjectName
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	return fs.WalkDir(templates.Templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walk error at %s: %w", path, err)
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(filepath.Base(path), ".") || filepath.Base(path) == "embed.go" {
			return nil
		}

		content, err := templates.Templates.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file %s failed: %w", path, err)
		}

		outputPath := filepath.Join(targetDir, strings.TrimSuffix(path, ".tmpl"))

		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return fmt.Errorf("mkdir %s failed: %w", filepath.Dir(outputPath), err)
		}

		if filepath.Ext(path) == ".tmpl" {
			tmpl, err := template.New(path).Parse(string(content))
			if err != nil {
				return fmt.Errorf("parse template %s failed: %w", path, err)
			}

			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, config); err != nil {
				return fmt.Errorf("execute template %s failed: %w", path, err)
			}

			if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
				return fmt.Errorf("write file %s failed: %w", outputPath, err)
			}
		} else {

			if err := os.WriteFile(outputPath, content, 0644); err != nil {
				return fmt.Errorf("write file %s failed: %w", outputPath, err)
			}
		}

		return nil
	})
}
