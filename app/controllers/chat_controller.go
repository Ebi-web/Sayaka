package controllers

import (
	"net/http"

	"Sayaka/gpt3"
	"Sayaka/utils"
)

type ChatRequest struct {
	Text string `json:"text"`
}

func Respond(_ http.ResponseWriter, r *http.Request) (int, string, error) {
	var req ChatRequest
	if err := utils.ParseRequestBody(r, &req); err != nil {
		return http.StatusBadRequest, "", err
	}

	text := req.Text
	res, err := gpt3.Chat(text)
	if err != nil {
		return http.StatusBadRequest, "", err
	}
	return http.StatusOK, res, nil
}
