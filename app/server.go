package main

import (
	"fmt"
	"net/http"

	"Sayaka/controllers"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/line/webhook", func(w http.ResponseWriter, r *http.Request) {
		if !middlewares.Verify(r) {
			fmt.Println("🚫Invalid signature")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		status, _ := controllers.ResLineWebhook(w, r)
		w.WriteHeader(status)
		return
	})
}
