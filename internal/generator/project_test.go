package generator

import (
	"os"
	"testing"
)

func TestGenerateProject(t *testing.T) {
	tmpDir := t.TempDir()
	
	// We want to generate the project INSIDE tmpDir/MyProject
	// Generate uses ProjectName as the target directory name in current working dir?
	// generator.go: targetDir := config.ProjectName
	// os.MkdirAll(targetDir, ...)
	// This will create it in the current directory (which is the test execution dir).
	// We should probably change the working directory or make Generate accept a base path.
	// But `Generate` implementation uses `os.MkdirAll(targetDir, 0755)`.
	// If I run this test, it will try to create "TestProj" in the package directory?
	// Tests run in the package directory.
	
	// To avoid cluttering, I should change to tmpDir.
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(tmpDir)

	config := ProjectConfig{
		ProjectName:  "TestProj",
		Module:       "example.com/testproj",
		Dependencies: []string{"dep1", "dep2"},
	}

	err := Generate(config)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Check if TestProj directory exists
	if _, err := os.Stat("TestProj"); os.IsNotExist(err) {
		t.Fatalf("Project directory not created")
	}

	// Check if embed.go exists (It SHOULD NOT)
	if _, err := os.Stat("TestProj/embed.go"); err == nil {
		t.Errorf("Bug: embed.go was generated in the output project")
	}

	// Check if go.mod exists
	if _, err := os.Stat("TestProj/go.mod"); os.IsNotExist(err) {
		t.Errorf("go.mod was not generated")
	}

    // Check if subdirectories are created (e.g. cmd/server)
    // Based on file list: cmd/server/main.go.tmpl
    if _, err := os.Stat("TestProj/cmd/server/main.go"); os.IsNotExist(err) {
         t.Errorf("cmd/server/main.go was not generated")
    }
}
