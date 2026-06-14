package sentiment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Analyzer interface {
	Analyze(
		title string,
		description string,
	) (string, error)
}

type OllamaAnalyzer struct {
	baseURL string
	model   string
	client  *http.Client
}

func (o *OllamaAnalyzer) Analyze(
	title string,
	description string,
) (string, error) {

	reqBody := generateRequest{
		Model: o.model,
		Prompt: buildPrompt(
			title,
			description,
		),
		Stream: false,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := o.client.Post(
		o.baseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result generateResponse

	if err := json.NewDecoder(resp.Body).
		Decode(&result); err != nil {

		return "", err
	}

	sentiment :=
		strings.TrimSpace(
			strings.ToUpper(
				result.Response,
			),
		)

	switch sentiment {

	case "POSITIVE":
		return sentiment, nil

	case "NEGATIVE":
		return sentiment, nil

	case "NEUTRAL":
		return sentiment, nil

	default:
		return "NEUTRAL", nil
	}
}

func NewOllamaAnalyzer() *OllamaAnalyzer {
	return &OllamaAnalyzer{
		baseURL: "http://localhost:11434",
		model:   "qwen3:0.6b",
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func buildPrompt(
	title string,
	description string,
) string {

	return fmt.Sprintf(`
Analyze sentiment.

Title:
%s

Description:
%s

Return only one word:

POSITIVE
NEGATIVE
NEUTRAL
`,
		title,
		description,
	)
}
