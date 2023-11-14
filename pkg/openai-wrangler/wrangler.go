package wrangler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

// GetJSONFromResponse is a high-level wrapper that contains the logic to
// extract the first content from a response and unmarshal it into the provided struct pointer
func GetJSONFromResponse(resp OpenAIResponse, result interface{}) error {
	content, err := extractFirstContent(resp)
	if err != nil {
		return err
	}

	return processContent(content, result)
}

// extractFirstContent extracts the first content from the response
func extractFirstContent(resp OpenAIResponse) (string, error) {
	if len(resp.Choices) == 0 {
		return "", errors.New("no choices found in response")
	}
	return resp.Choices[0].Message.Content, nil
}

// processContent processes the content of a response and unmarshals it into the provided struct pointer
func processContent(content string, result interface{}) error {
	cleanedContent, err := extractAndCleanJSON(content)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(cleanedContent), result)
	if err != nil {
		return err
	}

	// This is incredibly janky, but the easiest way to check if the JSON format was valid. If not the struct is set to nil values
	switch res := result.(type) {
	case *models.BeginnerProject:
		if res.Title == "" || len(res.Steps) == 0 || len(res.WatchOuts) == 0 || len(res.ExtraChallenges) == 0 || res.Objective == "" {
			return fmt.Errorf("required fields are missing or nil in BeginnerProject")
		}
	case *models.IntermediateProject:
		if res.Title == "" || len(res.Features) == 0 || len(res.Steps) == 0 || res.Objective == "" || res.Setup == "" || res.Testing == "" || res.Debugging == "" || len(res.Extras) == 0 {
			return fmt.Errorf("required fields are missing or nil in IntermediateProject")
		}
	case *models.ExperiencedProject:
		if res.Title == "" || res.Objective == "" || len(res.DetailedFeatures) == 0 || res.DevelopmentProcess.Setup == "" || len(res.DevelopmentProcess.Phases) == 0 || res.DevelopmentProcess.Testing == "" || res.DevelopmentProcess.Debugging == "" || len(res.Challenges) == 0 {
			return fmt.Errorf("required fields are missing or nil in ExperiencedProject")
		}

	default:
		return fmt.Errorf("unsupported project type")
	}
	return nil
}

// extractAndCleanJSON extracts the JSON content from the response and cleans it
func extractAndCleanJSON(content string) (string, error) {
	var startIndex, endIndex int
	var err error

	if strings.Contains(content, "```json") {
		startIndex, endIndex, err = findCodeBlockIndices(content, "```json", "```")
	} else if strings.Contains(content, "```") {
		startIndex, endIndex, err = findCodeBlockIndices(content, "```", "```")
	} else {
		// Hopefully the content is just JSON
		startIndex, endIndex = 0, len(content)
	}
	if err != nil {
		return "", err
	}

	cleanedContent := content[startIndex:endIndex]
	cleanedContent = strings.ReplaceAll(cleanedContent, "\\n", "")
	cleanedContent = strings.ReplaceAll(cleanedContent, "\\", "")

	return cleanedContent, nil
}

// findCodeBlockIndices finds the start and end indices of a code block in the content
func findCodeBlockIndices(content, startMarker, endMarker string) (int, int, error) {
	startIndex := strings.Index(content, startMarker)
	if startIndex == -1 {
		return 0, 0, errors.New("start marker not found")
	}
	// Adjust the start index to account for the length of the start marker
	endIndex := strings.Index(content[startIndex+len(startMarker):], endMarker)
	if endIndex == -1 {
		return -1, -1, errors.New("end marker not found")
	}
	// Adjust the end index to account for the length of the start marker
	endIndex += startIndex + len(startMarker)

	return startIndex + len(startMarker), endIndex, nil
}
