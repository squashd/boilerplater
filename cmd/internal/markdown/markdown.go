package markdown

import (
	"fmt"
	"strings"

	"github.com/SQUASHD/boilerplater/shared/models"
)

type Markdownable interface {
	ToMarkdown() string
}

type MarkdownGenerator struct{}

// GenerateMarkdownHeader generates a Markdown header.
func (m MarkdownGenerator) GenerateMarkdownHeader(level int, text string) string {
	return fmt.Sprintf("%s %s\n\n", strings.Repeat("#", level), text)
}

// GenerateMarkdownList generates a Markdown list.
func (m MarkdownGenerator) GenerateMarkdownList(items []string, ordered bool) string {
	var builder strings.Builder
	for i, item := range items {
		prefix := "- "
		if ordered {
			prefix = fmt.Sprintf("%d. ", i+1)
		}
		builder.WriteString(fmt.Sprintf("%s%s\n", prefix, item))
	}
	builder.WriteString("\n")
	return builder.String()
}

func (m MarkdownGenerator) GenerateMarkdownParagraph(text string) string {
	return fmt.Sprintf("%s\n\n", text)
}

func (m MarkdownGenerator) GenerateMarkdownCodeBlock(code, language string) string {
	return fmt.Sprintf("```%s\n%s\n```\n\n", language, code)
}

func (m MarkdownGenerator) GenerateMarkdownBlockquote(text string) string {
	return fmt.Sprintf("> %s\n\n", text)
}

func (m MarkdownGenerator) GenerateMarkdownImage(alt, url string) string {
	return fmt.Sprintf("![%s](%s)\n\n", alt, url)
}

func (m MarkdownGenerator) GenerateMarkdownLink(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

func (m MarkdownGenerator) ConvertStepsToMarkdown(steps []models.ProjectStep) string {
	var sb strings.Builder
	for i, step := range steps {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, step.Description))
		if step.Tips != nil && *step.Tips != "" {
			sb.WriteString(fmt.Sprintf("    - %s\n", *step.Tips))
		}
	}
	sb.WriteString("\n")
	return sb.String()
}

func (m MarkdownGenerator) ConverFeaturesToMarkdown(features []models.Feature) string {
	var sb strings.Builder
	for _, feature := range features {
		sb.WriteString(m.GenerateMarkdownHeader(3, feature.Name))
		sb.WriteString(m.GenerateMarkdownParagraph(feature.Description))
		if len(feature.Tips) > 0 {
			sb.WriteString(m.GenerateMarkdownHeader(4, "Tips"))
			sb.WriteString(m.GenerateMarkdownList(feature.Tips, false))
		}
	}
	return sb.String()
}

func (mg MarkdownGenerator) GenerateMarkdown(m Markdownable) string {
	return m.ToMarkdown()
}
