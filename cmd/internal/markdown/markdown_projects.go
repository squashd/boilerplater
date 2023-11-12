package markdown

import (
	"strings"

	"github.com/SQUASHD/boilerplater/shared/models"
)

type BeginnerProj struct {
	models.BeginnerProject
}

type IntermediateProj struct {
	models.IntermediateProject
}

type AdvancedProj struct {
	models.AdvancedProject
}

func (bp BeginnerProj) ToMarkdown() string {
	gen := MarkdownGenerator{}

	var sb strings.Builder
	// Title
	sb.WriteString(gen.GenerateMarkdownHeader(1, bp.Title))

	// Objective
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Objective"))
	sb.WriteString(gen.GenerateMarkdownParagraph(bp.Objective))

	// Steps
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Steps"))
	sb.WriteString(gen.ConvertStepsToMarkdown(bp.Steps))

	// Watchouts
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Watchouts"))
	sb.WriteString(gen.GenerateMarkdownList(bp.WatchOuts, false))

	// Extra Challenges
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Extra Challenges"))
	sb.WriteString(gen.GenerateMarkdownList(bp.ExtraChallenges, false))

	return sb.String()
}

func (ip IntermediateProj) ToMarkdown() string {
	gen := MarkdownGenerator{}

	var sb strings.Builder
	// Title
	sb.WriteString(gen.GenerateMarkdownHeader(1, ip.Title))

	// Objective
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Objective"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ip.Objective))

	// Features
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Features"))
	for _, feature := range ip.Features {
		sb.WriteString(gen.GenerateMarkdownHeader(3, feature.Name))
		sb.WriteString(gen.GenerateMarkdownParagraph(feature.Description))
		if len(feature.Tips) > 0 {
			sb.WriteString(gen.GenerateMarkdownHeader(4, "Tips"))
			sb.WriteString(gen.GenerateMarkdownList(feature.Tips, false))
		}
	}

	// Steps
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Steps"))
	sb.WriteString(gen.ConvertStepsToMarkdown(ip.Steps))
	// Setup
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Setup"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ip.Setup))

	// Testing
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Testing"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ip.Testing))

	// Debugging
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Debugging"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ip.Debugging))

	// Extras
	if len(ip.Extras) > 0 {
		sb.WriteString(gen.GenerateMarkdownHeader(2, "Extras"))
		sb.WriteString(gen.GenerateMarkdownList(ip.Extras, false))
	}

	return sb.String()
}

func (ap AdvancedProj) ToMarkdown() string {
	gen := MarkdownGenerator{}

	var sb strings.Builder
	// Title
	sb.WriteString(gen.GenerateMarkdownHeader(1, ap.Title))

	// Objective
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Objective"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ap.Objective))

	// Detailed Features
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Detailed Features"))
	for _, feature := range ap.DetailedFeatures {
		sb.WriteString(gen.GenerateMarkdownHeader(3, feature.Name))
		sb.WriteString(gen.GenerateMarkdownParagraph(feature.Description))
		if len(feature.ImplementationSteps) > 0 {
			sb.WriteString(gen.GenerateMarkdownHeader(4, "Implementation Steps"))
			sb.WriteString(gen.GenerateMarkdownList(feature.ImplementationSteps, false))
		}
	}

	// Development Process
	sb.WriteString(gen.GenerateMarkdownHeader(2, "Development Process"))
	// Setup
	sb.WriteString(gen.GenerateMarkdownHeader(3, "Setup"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ap.DevelopmentProcess.Setup))
	// Phases
	sb.WriteString(gen.GenerateMarkdownHeader(3, "Phases"))
	sb.WriteString(gen.GenerateMarkdownList(ap.DevelopmentProcess.Phases, true))
	// Testing
	sb.WriteString(gen.GenerateMarkdownHeader(3, "Testing"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ap.DevelopmentProcess.Testing))
	// Debugging
	sb.WriteString(gen.GenerateMarkdownHeader(3, "Debugging"))
	sb.WriteString(gen.GenerateMarkdownParagraph(ap.DevelopmentProcess.Debugging))

	// Challenges
	if len(ap.Challenges) > 0 {
		sb.WriteString(gen.GenerateMarkdownHeader(2, "Challenges"))
		sb.WriteString(gen.GenerateMarkdownList(ap.Challenges, false))
	}

	return sb.String()
}
