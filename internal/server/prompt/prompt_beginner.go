package prompt

import (
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

const (
	// TODO: Programmatically generate this JSON structure
	BeginnerProjectStruct = `
type BeginnerProject struct {
 Title string ` + "`json:\"title\"`" + `
 Objective string ` + "`json:\"objective\"`" + `
 Steps []ProjectStep ` + "`json:\"steps\"`" + `
 WatchOut []string ` + "`json:\"watchOuts\"`" + `
 ExtraChallenges []string ` + "`json:\"extraChallenges\"`" + `
}

type ProjectStep struct {
 Description string ` + "`json:\"description\"`" + `
 Tips *string ` + "`json:\"tips,omitempty\"`" + `
}
`
)

func generateBeginnerSystemPrompt(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString(BaseSystemPrompt)
	sb.WriteString("Your task is to create a project outline suitable for a beginner programmer, considering their overall proficiency and familiarity with the specified programming language. ")
	sb.WriteString("Focus on fundamental concepts and achievable challenges. Avoid suggesting advanced features or complex implementations that might overwhelm a beginner. ")
	sb.WriteString(SystemResponseAdherence)
	sb.WriteString(processStructString(BeginnerProjectStruct))
	sb.WriteString(" Focus on providing clear, step-by-step instructions and practical tips. ")
	sb.WriteString("Ensure the project is educational and engaging, helping the user solidify their understanding of basic programming concepts. ")
	sb.WriteString("If the project involves potentially complex tasks like API integration or web scraping, simplify these aspects and provide thorough guidance. ")
	sb.WriteString("Remember, the goal is to foster learning and confidence in programming, not to challenge the user with overly difficult tasks.")
	return sb.String()
}
