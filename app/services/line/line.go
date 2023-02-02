package line

import (
	"encoding/json"
	"net/url"
	"os"

	"Sayaka/services/line/messages"
	"Sayaka/utils"
)

const (
	baseEndpoint          = "https://api.line.me"
	replyEndpointSuffix   = "/v2/bot/message/reply/"
	verifyEndpointSuffix  = "/oauth2/v2.1/verify"
	profileEndpointSuffix = "/v2/profile"

	replyEndpoint   = baseEndpoint + replyEndpointSuffix
	verifyEndpoint  = baseEndpoint + verifyEndpointSuffix
	profileEndpoint = baseEndpoint + profileEndpointSuffix
)

type ReplyObject struct {
	ReplyToken string               `json:"replyToken"`
	Messages   []messages.TextReply `json:"messages"`
}

type AuthResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
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

func GetProfileByAccessToken(accessToken string) (AuthResponse, error) {
	method := "GET"
	headers := map[string]string{}

	params := url.Values{}
	params.Add("access_token", accessToken)
	ep := verifyEndpoint + "?" + params.Encode()

	_, err := utils.MakeRequest(method, ep, headers, nil)
	if err != nil {
		return AuthResponse{}, err
	}

	method = "GET"
	headers = map[string]string{
		"Authorization": "Bearer " + accessToken,
	}
	res, err := utils.MakeRequest(method, profileEndpoint, headers, nil)

	var authRes AuthResponse
	if err = json.Unmarshal(res, &authRes); err != nil {
		return AuthResponse{}, err
	}

	return authRes, nil
}
