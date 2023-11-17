package app

import (
	"os"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the application's state.
type App struct {
	AppRunning     bool
	RequestRunning bool // used for the spinner component
}

// NewApp creates a new instance of App.
func NewApp() *App {
	return &App{AppRunning: true}
}

// ExitAndCleanup exits the application and cleans up the terminal.
func (a *App) ExitAndCleanup(p *tea.Program) {
	if !a.AppRunning {
		err := p.ReleaseTerminal()
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

var SkillIntToString = map[int]string{
	0: "Beginner",
	1: "Intermediate",
	2: "Experienced",
}

var TargetOutcomeIntToString = map[int]string{
	0: "Learn",
	1: "Build",
	2: "Production-Ready",
}

// NewProjectConfig holds the configuration details for creating a new project.
type NewProjectConfig struct {
	Name               string
	Language           string
	Description        string
	OverallSkillLevel  int
	LanguageSkillLevel int
	TargetOutcome      int
}

// FieldValues converts the field values to more human-legible text
// This was mainly used during development.
func (p *NewProjectConfig) FieldValues() []string {
	return []string{
		p.Name,
		p.Language,
		p.Description,
		SkillIntToString[p.OverallSkillLevel],
		SkillIntToString[p.LanguageSkillLevel],
		TargetOutcomeIntToString[p.TargetOutcome],
	}
}

// ConvertProjConfigToRequest does what it says on the tin.
// Go enums are weird?
func ConvertProjConfigToRequest(p *NewProjectConfig) models.ProjectRequest {

	var profLevel models.ProficiencyLevel
	var expLevel models.ProficiencyLevel
	var desiredOutcome models.DesiredOutcome

	switch p.OverallSkillLevel {
	case 0:
		expLevel = models.Beginner
	case 1:
		expLevel = models.Intermediate
	case 2:
		expLevel = models.Experienced
	}

	switch p.LanguageSkillLevel {
	case 0:
		profLevel = models.Beginner
	case 1:
		profLevel = models.Intermediate
	case 2:
		profLevel = models.Experienced
	}

	switch p.TargetOutcome {
	case 0:
		desiredOutcome = models.Learning
	case 1:
		desiredOutcome = models.Portfolio
	case 2:
		desiredOutcome = models.Production
	}

	return models.ProjectRequest{
		ProjectName:     p.Name,
		Language:        p.Language,
		Description:     p.Description,
		LangProficiency: profLevel,
		Experience:      expLevel,
		DesiredOutcome:  desiredOutcome,
	}
}

// NewCommandContext is the context needed for executing the new command
// it encapsulates app state and the project config that is to be converted
// to a project request
type NewCommandContext struct {
	App     *App
	Project *NewProjectConfig
}

// NewNewCommandContext creates a new instance of NewCommandContext.
func NewNewCommandContext() *NewCommandContext {
	return &NewCommandContext{
		App:     NewApp(),
		Project: &NewProjectConfig{},
	}
}
