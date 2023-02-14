package gpt3

import (
	"encoding/json"
	"os"

	"Sayaka/utils"
)

const (
	baseURL = "https://api.openai.com"
)

type ChatRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	MaxTokens        int      `json:"max_tokens"`
	Temperature      float64  `json:"temperature"`
	PresencePenalty  float64  `json:"presence_penalty"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	TopP             float64  `json:"top_p"`
	Stop             []string `json:"stop"`
}

type ChatResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
	Usage   Usage     `json:"usage"`
}
type Choices struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func Chat(text string) (string, error) {
	path := "/v1/completions"
	url := baseURL + path
	method := "POST"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"),
	}

	req := &ChatRequest{
		Model:            "text-davinci-003",
		Prompt:           prompt(text),
		MaxTokens:        60,
		Temperature:      0.5,
		TopP:             1.0,
		FrequencyPenalty: 0.5,
		PresencePenalty:  0.0,
		Stop:             []string{"You:"},
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	res, err := utils.MakeRequest(method, url, headers, body)
	if err != nil {
		return "", err
	}

	var chatResponse ChatResponse
	if err = json.Unmarshal(res, &chatResponse); err != nil {
		return "", err
	}

	t := chatResponse.Choices[0].Text

	return t, nil
}

func prompt(text string) string {
	t := `The following is a conversation with a friend chatbot whose name is Sayaka. The friend is helpful, creative, clever, and very friendly.

You: What have you been up to?
Friend: Watching old movies.
You: Did you watch anything interesting?
Friend: Yeah, I watched an old classic called Casablanca. It was really good!

You: 
` + text + `
Friend: 
`
	return t
}
