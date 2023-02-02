package auth

import (
	"context"
	"encoding/json"
	"net/url"
	"os"

	"Sayaka/services/line"
	"Sayaka/utils"
	"github.com/pkg/errors"
)

type lineAuthRes struct {
	Iss     string   `json:"iss,omitempty"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud,omitempty"`
	Exp     int      `json:"exp,omitempty"`
	Iat     int      `json:"iat,omitempty"`
	Nonce   string   `json:"nonce,omitempty"`
	Amr     []string `json:"amr,omitempty"`
	Name    string   `json:"name"`
	Picture string   `json:"picture,omitempty"`
	Email   string   `json:"email,omitempty"`
}

type Res struct {
	LineUserID string `json:"line_user_id"`
}

func VerifyIDToken(_ context.Context, t string) (Res, error) {
	method := "POST"

	data := url.Values{}
	data.Set("id_token", t)
	data.Set("client_id", os.Getenv("AUTH_CHANNEL_ID"))

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := utils.MakeRequest(method, line.VerifyEndpoint, headers, data.Encode())
	if err != nil {
		return Res{}, errors.Wrap(err, "LINE ID token verification failed")
	}
	var lar lineAuthRes
	if err = json.Unmarshal(res, &lar); err != nil {
		return Res{}, errors.Wrap(err, "The LINE ID token was successfully validated but could not be mapped to a struct")
	}
	return Res{LineUserID: lar.Sub}, nil
}
