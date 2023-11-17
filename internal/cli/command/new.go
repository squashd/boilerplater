package command

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/SQUASHD/boilerplater/internal/cli/app"
	"github.com/SQUASHD/boilerplater/internal/cli/input"
	"github.com/SQUASHD/boilerplater/internal/cli/ui"
	"github.com/SQUASHD/boilerplater/internal/shared/httpclient"
	"github.com/SQUASHD/boilerplater/internal/shared/models"
	"github.com/SQUASHD/boilerplater/pkg/markdown"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Start a new project with Boilerplater",
	Long:  "Boilerplater will help you get started on your project by providing – hopefully – a clear structure with actionable steps to get you started.",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := app.NewNewCommandContext()

		ui.RenderLogo()
		gatherProjectDetails(ctx)

		var wg sync.WaitGroup

		ctx.App.RequestRunning = true
		wg.Add(2)

		go runSpinner(ctx, &wg)
		go handleProjectCreation(ctx, &wg)

		wg.Wait()
		ui.PrettyPrintChoice("Project created successfully!")
	},
}

// gatherProjectDetails is a wrapper function for the steps to gather all
// the information needed to generate a project outline
func gatherProjectDetails(ctx *app.NewCommandContext) {
	getProjectName(ctx)
	getLanguage(ctx)
	getSkillLevel(ctx)
	getProjectDescription(ctx)
	getLanguageProficiency(ctx)
	getTargetOutcome(ctx)
}

// runSpinner calls the UI spinner component and runs it in seperate go routine
// the project creation process can end the spinner component
func runSpinner(ctx *app.NewCommandContext, wg *sync.WaitGroup) {
	defer wg.Done()
	ui.RunSpinner(ctx.App)
}

// handleProjectCreation is a wrapper function to take care of processing
// the user input to generate a project outline
func handleProjectCreation(ctx *app.NewCommandContext, wg *sync.WaitGroup) {
	defer wg.Done()
	projReq := app.ConvertProjConfigToRequest(ctx.Project)

	project, err := fetchProjectFromServer(ctx, &projReq)
	if err != nil {
		handleCreationError(ctx, err)
		return
	}

	if err := setupProjectDirectory(ctx); err != nil {
		fmt.Println(err)
		return
	}

	if err := generateMarkdownFile(project); err != nil {
		fmt.Println(err)
		return
	}
}

// fetchProjectFromServer creates a new HTTP client and sends the project details
// for processing on the server. The model is determined based on the the experience
// level input.
func fetchProjectFromServer(ctx *app.NewCommandContext, projReq *models.ProjectRequest) (any, error) {
	client := httpclient.NewHTTPClient(2 * time.Minute)
	url := "http://localhost:8080/api/v1/templates"
	timeoutCtx, cancel := client.NewTimeoutContext(2 * time.Minute)
	defer cancel()

	model := determineModel(ctx)
	err := client.Post(timeoutCtx, url, projReq, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

// setupProjectDirectory creates the directory based on the project name input
func setupProjectDirectory(ctx *app.NewCommandContext) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	projectDir := path.Join(cwd, ctx.Project.Name)
	if err = os.Mkdir(projectDir, 0755); err != nil {
		return err
	}
	return os.Chdir(projectDir)
}

// generateMarkdownFile takes the project model and does a runtime check
// to see if the equivalent struct has implemented the markdownable interface
func generateMarkdownFile(model any) error {
	mdModel, err := convertToMarkdownable(model)
	if err != nil {
		return err
	}

	mdGen := markdown.MarkdownGenerator{}
	md := mdGen.GenerateMarkdown(mdModel)
	return os.WriteFile("PROJECT.md", []byte(md), 0644)
}

func convertToMarkdownable(model interface{}) (markdown.Markdownable, error) {
	switch m := model.(type) {
	case *models.BeginnerProject:
		return &markdown.BeginnerProj{BeginnerProject: *m}, nil
	case *models.IntermediateProject:
		return &markdown.IntermediateProj{IntermediateProject: *m}, nil
	case *models.ExperiencedProject:
		return &markdown.ExperiencedProj{ExperiencedProject: *m}, nil
	default:
		return nil, fmt.Errorf("unknown model type")
	}
}

// handleCreationError is a simple wrapper to handle errors that can occur during creation,
// ensuring processes end and the terminal is released
func handleCreationError(ctx *app.NewCommandContext, err error) {
	fmt.Println(err)
	ctx.App.RequestRunning = false
	ctx.App.AppRunning = false
}

// determineModel takes the input int setst he correct model
func determineModel(ctx *app.NewCommandContext) interface{} {
	var model interface{}

	switch ctx.Project.OverallSkillLevel {
	case 0:
		model = &models.BeginnerProject{}
	case 1:
		model = &models.IntermediateProject{}
	case 2:
		model = &models.ExperiencedProject{}
	}

	return model
}

// getProjectName gets the input for the directory name that the project is to
// be created in. The input is validated for alphanumeric values
func getProjectName(ctx *app.NewCommandContext) {
	err := input.NewTextInputInteraction(
		ctx.App,
		"What's the name of your project?",
		"project-name",
		20,
		20,
		&ctx.Project.Name).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(ctx.Project.Name)
}

// getSkillLevel is a multichoice prompt for self-described skill level
// further options are limited based on the initial input
func getSkillLevel(ctx *app.NewCommandContext) {
	err := input.NewChoiceInteraction(
		ctx.App,
		"What kind of developer are you?",
		input.SkillLevelOptions(),
		&ctx.Project.OverallSkillLevel).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(app.SkillIntToString[ctx.Project.OverallSkillLevel])
}

// getLanguage gets the preferred language for the project. This is later
// used to prompt the user for their proficiency
func getLanguage(ctx *app.NewCommandContext) {
	err := input.NewTextInputInteraction(
		ctx.App,
		"What language do you want to write this project in?",
		"python",
		20,
		20,
		&ctx.Project.Language).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(ctx.Project.Language)
}

// getProjectDescription is prompt for a preferably short description of their project
func getProjectDescription(ctx *app.NewCommandContext) {
	err := input.NewTextInputInteraction(
		ctx.App,
		"Describe your project - shorter is better",
		"a to-do list app",
		140,
		140,
		&ctx.Project.Description).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(ctx.Project.Description)
}

// getLanguageProficiency gets the self-evaluated skill level in a language
// user's with a self-evaluated overall skill level are determined to be
// beginner's in their determined language, skipping the prompting phase
// while intermediate and experienced may have a "lower" level in their chosen language
func getLanguageProficiency(ctx *app.NewCommandContext) {

	if ctx.Project.OverallSkillLevel == 0 {
		ctx.Project.LanguageSkillLevel = 0
		return
	}

	choices := input.FilterChoices(ctx.Project.OverallSkillLevel, input.LanguageSkillLevelOptions())

	formattedLanguagePrompt := fmt.Sprint("What is your experience with ", ctx.Project.Language, "?")
	err := input.NewChoiceInteraction(
		ctx.App,
		formattedLanguagePrompt,
		choices,
		&ctx.Project.LanguageSkillLevel).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(app.SkillIntToString[ctx.Project.LanguageSkillLevel])
}

// getTargetOutcome is more of a vibe check that hopefully prompts the AI to be
// more or less stringent in the requirements for the project
func getTargetOutcome(ctx *app.NewCommandContext) {

	if ctx.Project.OverallSkillLevel == 0 {
		ctx.Project.TargetOutcome = 0
		return
	}

	err := input.NewChoiceInteraction(
		ctx.App,
		"What type of project is this?",
		input.TargetOutcomeOptions(),
		&ctx.Project.TargetOutcome).
		Execute()

	if err != nil {
		fmt.Println(err)
	}
	ui.PrettyPrintChoice(app.TargetOutcomeIntToString[ctx.Project.TargetOutcome])
}
