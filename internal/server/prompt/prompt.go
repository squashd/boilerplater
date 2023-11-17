package prompt

import (
	"fmt"
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

// Some thoughtfully crafted prompts that - hopefully - generates more consistent
// output from the AI
const (
	BaseSystemPrompt        = "You are an experienced developer and educator. "
	SystemResponseAdherence = "Your response must adhere to this JSON structure: "
)

// GenerateSystemPrompt is a switch to generate the system prompt
func GenerateSystemPrompt(req *models.ProjectRequest) string {
	switch req.Experience {
	case 0:
		return generateBeginnerSystemPrompt(req)
	case 1:
		return generateIntermidateSystemPrompt(req)
	case 2:
		return generateExperiencedSystemPrompt(req)
	default:
		return ""
	}
}

// GenerateUserPrompt is a wrapper to generate the user prompt
func GenerateUserPrompt(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString(generateSkillStatement(req))
	sb.WriteString(generateProjectDescription(req))
	sb.WriteString(generateLanguageProficiencyStatement(req))
	sb.WriteString(generateBasicFamiliarityStatement(req))
	sb.WriteString(generateDesiredOutcome(req))
	fmt.Printf("Generated user prompt: %s\n", sb.String())
	return sb.String()
}

func generateSkillStatement(req *models.ProjectRequest) string {
	switch req.Experience {
	case 0:
		return "I am a beginner programmer."
	case 1:
		return "I am an intermediate programmer."
	case 2:
		return "I am an experienced programmer."
	}
	return ""
}

// generateProjectDescription combines the project description with the language chosen
// TODO: implement 'Surprise Me' or 'Any lanaguage' choice
func generateProjectDescription(req *models.ProjectRequest) string {
	return "I want to create " + req.Description + " using " + req.Language + ". "
}

// generateLanguageProficiencyStatement hopefully makes the AI elaborate or focus
// on unique language details
func generateLanguageProficiencyStatement(req *models.ProjectRequest) string {
	skill := strings.ToLower(req.Experience.String())
	if req.LangProficiency < req.Experience {
		return "Despite being a " + skill + " programmer, I am " + req.LangProficiency.String() + " with regards to " + req.Language + ". "
	}
	return ""
}

// generateBasicFamiliarityStatement tells the AI that despite being a self-professed
// beginner, they are already familiar with the basic, which reduced the focus
// on basic syntax
func generateBasicFamiliarityStatement(req *models.ProjectRequest) string {
	if req.Experience == models.Beginner {
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

// processStructString was made because I had some issues removing tabs and newlines
// from the string represenation of the structs and I though f*** it
// also makes the request slightly more compact
func processStructString(str string) string {
	cleaned := strings.ReplaceAll(str, "\t", "")
	cleaned = strings.ReplaceAll(cleaned, "\n", "")
	return cleaned
}
