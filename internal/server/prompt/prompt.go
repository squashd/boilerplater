package prompt

import (
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

const (
	BaseSystemPrompt        = "You are an experienced developer and educator. "
	SystemResponseAdherence = "Your response must adhere to this JSON structure: "
)

func GenerateSystemPrompt(req *models.ProjectRequest) string {
	switch req.ProficiencyLevel.String() {
	case "Beginner":
		return generateBeginnerSystemPrompt(req)
	case "Intermediate":
		return generateIntermidateSystemPrompt(req)
	case "Experienced":
		return generateExperiencedSystemPrompt(req)
	default:
		return ""
	}
}

func GenerateUserPrompt(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString(generateProgrammerProficiencyStatement(req))
	sb.WriteString(generateProjectDescription(req))
	sb.WriteString(generateLanguageProficiencyStatement(req))
	sb.WriteString(generateBasicFamiliarityStatement(req))
	sb.WriteString(generateDesiredOutcome(req))
	return sb.String()
}

func generateProgrammerProficiencyStatement(req *models.ProjectRequest) string {
	return "I am a " + req.ProficiencyLevel.String() + " programmer. "
}

func generateProjectDescription(req *models.ProjectRequest) string {
	return "I want to create " + req.Description + " using " + req.Language + ". "
}

func generateLanguageProficiencyStatement(req *models.ProjectRequest) string {
	if req.LanguageProficiency < req.ProficiencyLevel {
		return "Despite being a " + req.ProficiencyLevel.String() + " programmer, I am " + req.LanguageProficiency.String() + " with regards to " + req.Language + ". "
	}
	return ""
}

func generateBasicFamiliarityStatement(req *models.ProjectRequest) string {
	if req.ProficiencyLevel == models.Beginner {
		return "I am already familiar with coding basics. "
	}
	return ""
}

// generateDesiredOutcome generates a human legible statement about the desired outcome of the project
// based on the desired outcome enum
func generateDesiredOutcome(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString("I want this to be ")
	switch req.DesiredOutcome {
	case models.Learning:
		sb.WriteString("a learning experience")
	case models.Portfolio:
		sb.WriteString("a portfolio project")
	case models.Production:
		sb.WriteString("a production-ready project")
	}
	sb.WriteString(".")
	return sb.String()
}

func processStructString(str string) string {
	cleaned := strings.ReplaceAll(str, "\t", "")
	cleaned = strings.ReplaceAll(cleaned, "\n", "")
	return cleaned
}
