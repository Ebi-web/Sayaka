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
	"github.com/rs/cors"
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

	sql := db.NewPostgreSQL(databaseDatasource)
	dbConnection, err := sql.Open()
	if err != nil {
		return fmt.Errorf("🚫failed db init. %s", err)
	}
	s.db = dbConnection

	s.router = s.Route()
	return nil
}

func (s *Server) Route() *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.NewAliceChain(corsMiddleware.Handler)

	m := middlewares.NewValidateSignatureMiddleware()
	webhookChain := commonChain.Append(m.Handle)

	am := middlewares.NewAuth(s.db)
	authChain := commonChain.Append(am.Handle)

	webhookHandler := handler.NewWebhookHandler()
	userHandler := handler.NewUserHandler(s.db)
	flashCardHandler := handler.NewFlashCardHandler(s.db)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("200"))
	})

	r.Methods(http.MethodPost, http.MethodOptions).Path("/line/webhook").Handler(webhookChain.Then(AppHandler{webhookHandler.ResLineWebhook}))

	r.Methods(http.MethodPost, http.MethodOptions).Path("/users/register").Handler(commonChain.Then(AppHandler{userHandler.Register}))

	r.Methods(http.MethodPost, http.MethodOptions).Path("/flash-cards").Handler(authChain.Then(AppHandler{flashCardHandler.Create}))
	r.Methods(http.MethodGet, http.MethodOptions).Path("/flash-cards").Handler(authChain.Then(AppHandler{flashCardHandler.Index}))

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
