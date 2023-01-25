package middlewares

import "net/http"

type Middleware interface {
	Handle(h http.Handler) http.Handler
}
