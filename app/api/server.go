package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"Sayaka/api/handler"
	"Sayaka/db"
	"Sayaka/lib/alice"
	"Sayaka/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db     *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(databaseDatasource string) error {
	fmt.Println("ℹ️Server Initialising...")

	sql := db.NewMySQL(databaseDatasource)
	dbConnection, err := sql.Open()
	if err != nil {
		return fmt.Errorf("🚫failed db init. %s", err)
	}
	s.db = dbConnection

	s.router = s.Route()
	return nil
}

func (s *Server) Route() *mux.Router {
	commonChain := alice.NewAliceChain()

	m := middlewares.NewValidateSignatureMiddleware()
	webhookChain := commonChain.Append(m.Handle)

	webhookHandler := handler.NewWebhookHandler()
	userHandler := handler.NewUserHandler(s.db)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("200"))
	})

	r.Methods(http.MethodPost, http.MethodOptions).Path("/line/webhook").Handler(webhookChain.Then(AppHandler{webhookHandler.ResLineWebhook}))

	r.Methods(http.MethodPost, http.MethodOptions).Path("/users").Handler(commonChain.Then(AppHandler{userHandler.Create}))

	return r
}

func (s *Server) Run(port int) {
	log.Printf("🍺Listening on port %d", port)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	); err != nil {
		panic(err)
	}
}
