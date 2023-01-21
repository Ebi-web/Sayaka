package controllers

import (
	"net/http"

	"Sayaka/services/gpt3"
	"Sayaka/services/line"
	"Sayaka/utils"
)

func ResLineWebhook(_ http.ResponseWriter, r *http.Request) (int, error) {
	var webhookRequest line.WebhookRequest
	if err := utils.ParseRequestBody(r, &webhookRequest); err != nil {
		return http.StatusBadRequest, err
	}

	events := webhookRequest.Events
	for k := range events {
		if err := eventHandler(&events[k]); err != nil {
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}

func eventHandler(e *line.Event) error {
	if e.Type != "message" || e.Message.Type != "text" {
		return nil
	}
	//	ChatGPTにテキストを渡す
	text, err := gpt3.Chat(e.Message.Text)
	if err != nil {
		return err
	}
	// LINEの返信APIを叩く
	if err = line.Reply(e.ReplyToken, text); err != nil {
		return err
	}
	return nil
}
