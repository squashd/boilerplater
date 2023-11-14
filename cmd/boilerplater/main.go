package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/SQUASHD/boilerplater/internal/shared/httpclient"
	"github.com/SQUASHD/boilerplater/internal/shared/models"
	"github.com/SQUASHD/boilerplater/pkg/boilerplate"
	"github.com/SQUASHD/boilerplater/pkg/markdown"
	"github.com/SQUASHD/boilerplater/pkg/projectgen"
)

func main() {
	type ProjectData struct {
		ProjectStructure    models.ProjectStructure      `json:"projectStructure"`
		FunctionBoilerplate []models.FunctionBoilerplate `json:"functionBoilerplate"`
		AdvancedProject     markdown.ExperiencedProj     `json:"advancedProject"`
	}

	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	url := flag.String("url", "", "URL to fetch JSON from")
	flag.Parse()

	fmt.Println(os.Args)
	if *url == "" {
		log.Fatal("No URL provided")
	}

	var projectData ProjectData
	httpClient := httpclient.NewHTTPClient(2 * time.Minute)
	if err = httpClient.Get(nil, *url, &projectData); err != nil {
		log.Fatalf("Error fetching JSON: %v", err)
	}

	mdgen := markdown.MarkdownGenerator{}

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
