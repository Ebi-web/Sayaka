package alice

import (
	"github.com/justinas/alice"
)

func NewAliceChain(middlewares ...alice.Constructor) alice.Chain {
	return alice.New(newConstructors(middlewares...)...)
}

func newConstructors(middlewares ...alice.Constructor) []alice.Constructor {
	var constructors []alice.Constructor
	for k := range middlewares {
		constructors = append(constructors, middlewares[k])
	}
	return constructors
}
