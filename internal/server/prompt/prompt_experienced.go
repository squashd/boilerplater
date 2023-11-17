package prompt

import (
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

const (
	ExperiencedProjectStruct = `
type ExperiencedProject struct {
 Title string ` + "`json:\"title\"`" + `
 Objective string ` + "`json:\"objective\"`" + `
 DetailedFeatures []DetailedFeature ` + "`json:\"detailedFeatures\"`" + `
 DevelopmentProcess DevelopmentProcess ` + "`json:\"developmentProcess\"`" + `
 Challenges []string ` + "`json:\"challenges\"`" + `
}

type DetailedFeature struct {
 Name string ` + "`json:\"name\"`" + `
 Description string ` + "`json:\"description\"`" + `
 ImplementationSteps []string ` + "`json:\"implementationSteps\"`" + `
}

type DevelopmentProcess struct {
 Setup string ` + "`json:\"setup\"`" + `
 Phases []string ` + "`json:\"phases\"`" + `
 Testing string ` + "`json:\"testing\"`" + `
 Debugging string ` + "`json:\"debugging\"`" + `
}
`
)

// generateExperiencedSystemPrompt is made by someone who at best can be described
// as 'intermediate'.
// whether or not it actually produces something 'intellectually stimulating' remains to be seen
func generateExperiencedSystemPrompt(req *models.ProjectRequest) string {
	var sb strings.Builder
	sb.WriteString(BaseSystemPrompt)
	sb.WriteString("Develop a project outline for an experienced programmer that challenges their advanced skills in programming, software design, and architecture. ")
	sb.WriteString("The project should focus on complex concepts such as performance optimization, scalability, security, and advanced design patterns. ")
	sb.WriteString(SystemResponseAdherence)
	sb.WriteString(processStructString(ExperiencedProjectStruct))
	sb.WriteString("The project should be intellectually stimulating and encourage the exploration of advanced programming techniques. ")
	sb.WriteString("It should include the integration of complex external APIs or libraries, and the use of sophisticated testing and debugging methods. ")
	sb.WriteString("The aim is to foster a deep understanding and mastery of advanced programming concepts, while introducing complex tasks and industry-standard best practices. ")
	return sb.String()
}
