package controllers

import (
	"net/http"

	"Sayaka/services/gpt3"
	"Sayaka/services/line"
	"Sayaka/utils"
)

type WebhookRequest struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

type Event struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"message"`
	WebhookEventID  string `json:"webhookEventId"`
	DeliveryContext struct {
		IsRedelivery bool `json:"isRedelivery"`
	} `json:"deliveryContext"`
	Timestamp int64 `json:"timestamp"`
	Source    struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
	ReplyToken string `json:"replyToken"`
	Mode       string `json:"mode"`
}

func ResLineWebhook(_ http.ResponseWriter, r *http.Request) (int, error) {
	var webhookRequest WebhookRequest
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

func eventHandler(e *Event) error {
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
