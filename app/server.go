package main

import (
	"fmt"
	"net/http"

	"Sayaka/controllers"
	"Sayaka/lib/alice"
	"Sayaka/middlewares"
)

const PORT = "8080"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	s.Route()
	fmt.Println("🍺Starting server on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

func (s *Server) Route() {
	m := middlewares.NewValidateSignatureMiddleware()
	webhookChain := alice.NewAliceChain(m.Handle)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.Handle("/line/webhook", webhookChain.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status, _ := controllers.ResLineWebhook(w, r)
		w.WriteHeader(status)
		return
	})))
}
