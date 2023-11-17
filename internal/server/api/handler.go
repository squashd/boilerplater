package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/SQUASHD/boilerplater/internal/server/prompt"
	"github.com/SQUASHD/boilerplater/internal/shared/models"
	wrangler "github.com/SQUASHD/boilerplater/pkg/openai-wrangler"
)

// Because the request to the open AI API is only made here I've also defined it here
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// templatesHandlerV1 takes a request based on the shard project request struct
// sends it to OpenAI, recieves the response and sanitizes it for code block strings
// before sending it back to the user
func (s *Server) templatesHandlerV1(w http.ResponseWriter, r *http.Request) {
	req := &models.ProjectRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userPrompt := prompt.GenerateUserPrompt(req)
	sytemPrompt := prompt.GenerateSystemPrompt(req)

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
	aiReq.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI_API_KEY"))
	aiReq.Header.Set("OpenAI-Organization", os.Getenv("OPEN_AI_ORG_ID"))

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

	// The project outline is different depending on the experience level
	var model interface{}
	switch req.Experience {
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

	jsonRes, _ := json.Marshal(model)
	fmt.Printf("Project JSON: %+v\n", string(jsonRes))

	RespondWithJSON(w, http.StatusOK, model)
}
