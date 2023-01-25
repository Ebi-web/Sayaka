package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"os"
)

type ValidateSignatureMiddleware struct {
}

var _ Middleware = &ValidateSignatureMiddleware{}

func NewValidateSignatureMiddleware() *ValidateSignatureMiddleware {
	return &ValidateSignatureMiddleware{}
}

func (v *ValidateSignatureMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		body, err := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(body))
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(req.Header.Get("x-line-signature"))
		if err != nil {
			http.Error(w, "Failed to decode signature", http.StatusBadRequest)
			return
		}
		hash := hmac.New(sha256.New, []byte(os.Getenv("LINE_CHANNEL_SECRET")))
		hash.Write(body)

		if !hmac.Equal(decoded, hash.Sum(nil)) {
			http.Error(w, "Invalid signature", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, req)
	})
}
