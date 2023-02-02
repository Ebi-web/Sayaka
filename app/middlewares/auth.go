package middlewares

import (
	"log"
	"net/http"

	"Sayaka/auth"
	"Sayaka/domain/repository"
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
		verified, err := auth.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			utils.RespondErrorJson(w, http.StatusBadRequest, err)
			return
		}
		usr, err := repository.FindUserByLINEID(a.db, verified.LineUserID)
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
