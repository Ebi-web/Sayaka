package line

import (
	"encoding/json"
	"os"

	"Sayaka/services/line/messages"
	"Sayaka/utils"
)

const (
	baseEndpoint        = "https://api.line.me/v2/bot"
	replyEndpointSuffix = "/message/reply/"

	replyEndpoint = baseEndpoint + replyEndpointSuffix
)

type ReplyObject struct {
	ReplyToken string               `json:"replyToken"`
	Messages   []messages.TextReply `json:"messages"`
}

func Reply(token string, text string) error {
	textReplyObj := messages.NewTextReply(text)
	obj := &ReplyObject{
		ReplyToken: token,
		Messages:   []messages.TextReply{textReplyObj.Struct()},
	}
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	method := "POST"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	}
	_, err = utils.MakeRequest(method, replyEndpoint, headers, body)
	if err != nil {
		return err
	}
	return nil
}
