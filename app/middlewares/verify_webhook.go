package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"os"
)

func Verify(req *http.Request) bool {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return false
	}
	decoded, err := base64.StdEncoding.DecodeString(req.Header.Get("x-line-signature"))
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(os.Getenv("LINE_CHANNEL_SECRET")))
	hash.Write(body)

	return hmac.Equal(decoded, hash.Sum(nil))
}
