package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/SQUASHD/boilerplater/internal/server/prompt"
	"github.com/SQUASHD/boilerplater/internal/shared/models"
	wrangler "github.com/SQUASHD/boilerplater/pkg/openai-wrangler"
)

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (s *Server) templatesHandlerV1(w http.ResponseWriter, r *http.Request) {
	req := &models.ProjectRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	sytemPrompt := prompt.GenerateSystemPrompt(req)
	userPrompt := prompt.GenerateUserPrompt(req)

	payload := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user", Content: userPrompt},
			{Role: "system", Content: sytemPrompt},
		},
		Temperature: 0.7,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error marshalling request")
		return
	}

	aiReq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error connecting to AI server")
		return
	}

	aiReq.Header.Set("Content-Type", "application/json")
	aiReq.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_APIKEY"))
	aiReq.Header.Set("OpenAI-Organization", os.Getenv("OPENAI_ORGID"))

	client := &http.Client{}
	resp, err := client.Do(aiReq)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error connecting to AI server")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error reading response")
		return
	}

	var aiResponse wrangler.OpenAIResponse
	err = json.Unmarshal(body, &aiResponse)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error unmarshalling response")
		return
	}

	var model interface{}
	switch req.ProficiencyLevel {
	case 0:
		model = &models.BeginnerProject{}
	case 1:
		model = &models.IntermediateProject{}
	case 2:
		model = &models.ExperiencedProject{}
	}

	// Make sure it's a pointer to a struct and not the interface itself
	if err = wrangler.GetJSONFromResponse(aiResponse, model); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error processing response")
		return
	}

	RespondWithJSON(w, http.StatusOK, model)
}

func (s *Server) mockHandler(w http.ResponseWriter, r *http.Request) {

	type ProjectData struct {
		ProjectStructure    models.ProjectStructure      `json:"projectStructure"`
		FunctionBoilerplate []models.FunctionBoilerplate `json:"functionBoilerplate"`
		AdvancedProject     models.ExperiencedProject    `json:"advancedProject"`
	}

	data, err := os.ReadFile("/Users/hjartland/repos/personal-projects/boilerplater/internal/server/api/mock.json")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error reading mock data")
		return
	}

	var projectData ProjectData
	if err = json.Unmarshal(data, &projectData); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error unmarshalling mock data")
		return
	}

	RespondWithJSON(w, http.StatusOK, projectData)
}

func (s *Server) testPromptHandlerV1(w http.ResponseWriter, r *http.Request) {
	req := &models.ProjectRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	sytemPrompt := prompt.GenerateSystemPrompt(req)
	userPrompt := prompt.GenerateUserPrompt(req)

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"systemPrompt": sytemPrompt,
		"userPrompt":   userPrompt,
	})
}
