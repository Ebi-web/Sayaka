package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"Sayaka/domain/repository"
	"Sayaka/services/line"
	"Sayaka/utils"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	bearer = "Bearer"
)

type Auth struct {
	db *sqlx.DB
}

var _ Middleware = &Auth{}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken, err := getTokenFromHeader(r)
		if err != nil {
			utils.RespondErrorJson(w, http.StatusUnauthorized, err)
			return
		}
		verified, err := verifyIDToken(r.Context(), idToken)
		if err != nil {
			utils.RespondErrorJson(w, http.StatusBadRequest, err)
			return
		}
		usr, err := repository.FindUserByLINEID(a.db, verified.LINEUserID)
		if err != nil {
			log.Print(err.Error())
			utils.RespondErrorJson(w, http.StatusInternalServerError, err)
			return
		}
		ctx := utils.SetUserToContext(r.Context(), usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenFromHeader(req *http.Request) (string, error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header not found")
	}

	l := len(bearer)
	if len(header) > l+1 && header[:l] == bearer {
		return header[l+1:], nil
	}

	return "", errors.New("authorization header format must be 'Bearer {token}'")
}

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

type VerifiedLINEUser struct {
	LINEUserID string `json:"line_user_id"`
}

func verifyIDToken(_ context.Context, t string) (VerifiedLINEUser, error) {
	method := "POST"

	data := url.Values{}
	data.Set("id_token", t)
	data.Set("client_id", os.Getenv("AUTH_CHANNEL_ID"))

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := utils.MakeRequest(method, line.VerifyEndpoint, headers, data.Encode())
	if err != nil {
		return VerifiedLINEUser{}, errors.Wrap(err, "LINE ID token verification failed")
	}
	var authRes lineAuthRes
	if err = json.Unmarshal(res, &authRes); err != nil {
		return VerifiedLINEUser{}, errors.Wrap(err, "The LINE ID token was successfully validated but could not be mapped to a struct")
	}
	return VerifiedLINEUser{LINEUserID: authRes.Sub}, nil
}
