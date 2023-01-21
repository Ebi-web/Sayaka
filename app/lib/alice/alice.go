package alice

import (
	"net/http"

	"github.com/justinas/alice"
)

func NewAlice(handler http.Handler, middlewares ...func(h http.Handler) http.Handler) http.Handler {
	return alice.New(newConstructors(middlewares...)...).Then(handler)
}

func newConstructors(middlewares ...func(h http.Handler) http.Handler) []alice.Constructor {
	var constructors []alice.Constructor
	for k := range middlewares {
		constructors = append(constructors, newConstructor(middlewares[k]))
	}
	return constructors
}

func newConstructor(middleware func(h http.Handler) http.Handler) alice.Constructor {
	return middleware
}
