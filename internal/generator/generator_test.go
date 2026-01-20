package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateComponent(t *testing.T) {
	// Create a temporary directory for output
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test_handler.go")

	config := ComponentConfig{
		Name:         "TestUser",
		TemplateName: "handler",
		OutputPath:   outputPath,
	}

	err := GenerateComponent(config)
	if err != nil {
		t.Fatalf("GenerateComponent failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("File was not created at %s", outputPath)
	}

	// Verify content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	expectedStr := "func (s *Server) handleTestUser(w http.ResponseWriter, r *http.Request)"
	if !strings.Contains(string(content), expectedStr) {
		t.Errorf("File content does not match.\nExpected to contain: %s\nGot:\n%s", expectedStr, string(content))
	}
}
