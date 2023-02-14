package handler

import (
	"fmt"
	"net/http"

	"Sayaka/services/gpt3"
	"Sayaka/services/line"
	"Sayaka/services/line/webhook"
	"Sayaka/utils"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) ResLineWebhook(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var webhookRequest webhook.Request
	if err := utils.ParseRequestBody(r, &webhookRequest); err != nil {
		fmt.Println("Error: ", err)
		return http.StatusBadRequest, nil, err
	}

	events := webhookRequest.Events
	for k := range events {
		if err := eventHandler(&events[k]); err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}
	return http.StatusOK, nil, nil
}

func eventHandler(e *webhook.Event) error {
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
