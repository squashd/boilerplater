package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/SQUASHD/boilerplater/cmd/internal/boilerplate"
	"github.com/SQUASHD/boilerplater/cmd/internal/markdown"
	"github.com/SQUASHD/boilerplater/cmd/internal/projectgen"
	"github.com/SQUASHD/boilerplater/shared/models"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: boilerplater -i <path>")
	}
	if os.Args[1] != "-i" {
		log.Fatal("Usage: boilerplater -i <path>")
	}

	type ProjectData struct {
		ProjectStructure    models.ProjectStructure      `json:"projectStructure"`
		FunctionBoilerplate []models.FunctionBoilerplate `json:"functionBoilerplate"`
		AdvancedProject     markdown.AdvancedProj        `json:"advancedProject"`
	}

	inFilePath := os.Args[2]
	mdgen := markdown.MarkdownGenerator{}

	// projectDir is the directory containing the file at inFilePath
	projectDir := filepath.Dir(inFilePath)
	// Read the JSON file
	absPath, err := filepath.Abs(inFilePath)
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Decode the JSON file
	var projectData ProjectData
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&projectData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	if err = projectgen.GenerateProjectStructure(projectData.ProjectStructure); err != nil {
		log.Fatalf("Error generating project structure: %v", err)
	}
	// Generate the project structure
	boilerplate.GenerateFunctionBoilerplate(projectData.FunctionBoilerplate)

	// Write the md file based on mdgen output
	mdFilePath := filepath.Join(projectDir, "README.md")
	_, err = os.Create(mdFilePath)
	if err != nil {
		log.Fatalf("Error creating md file: %v", err)
	}
	err = os.WriteFile(mdFilePath, []byte(mdgen.GenerateMarkdown(projectData.AdvancedProject)), 0644)
	if err != nil {
		log.Fatalf("Error writing md file: %v", err)
	}

}
