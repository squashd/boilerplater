package prompt

import (
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

const (
	// TODO: Programmatically generate this JSON structure
	IntermediateProjectStruct = `
type IntermediateProject struct {
 Title string ` + "`json:\"title\"`" + `
 Objective string ` + "`json:\"objective\"`" + `
 Features []Feature ` + "`json:\"features\"`" + `
 Steps []ProjectStep ` + "`json:\"steps\"`" + `
 Setup string ` + "`json:\"setup\"`" + `
 Testing string ` + "`json:\"testing\"`" + `
 Debugging string ` + "`json:\"debugging\"`" + `
 Extras []string ` + "`json:\"extras\"`" + `
} 

type Feature struct {
 Name string ` + "`json:\"name\"`" + `
 Description string ` + "`json:\"description\"`" + `
 Tips []string ` + "`json:\"tips\"`" + `
} 

type ProjectStep struct {
 Description string ` + "`json:\"description\"`" + `
 Tips *string ` + "`json:\"tips,omitempty\"`" + `
}
`
)

// generateIntermidateSystemPrompt tries to focus on some higher level implementation details
func generateIntermidateSystemPrompt(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString(BaseSystemPrompt)
	sb.WriteString("Your task is to create a project outline suitable for an intermediate programmer, considering their overall proficiency in programming. ")
	sb.WriteString("Encourage the development of features that require a good understanding of programming concepts and introduce aspects of software design and architecture. ")
	sb.WriteString("Emphasize the importance of version control, comprehensive documentation, and best practices in coding. ")
	sb.WriteString(SystemResponseAdherence)
	sb.WriteString(processStructString(IntermediateProjectStruct))
	sb.WriteString("Focus on providing clear, step-by-step instructions and practical tips that align with the intermediate programmer's learning curve. ")
	sb.WriteString("Ensure the project is educational and engaging, helping the user to deepen their understanding of programming concepts. ")
	sb.WriteString("Encourage the integration of external APIs or libraries, and the use of testing and debugging techniques. ")
	sb.WriteString("Remember, the goal is to foster learning and confidence in programming, while introducing more complex tasks and best practices.")
	return sb.String()
}
