package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goldsmithb/spotted_lantern_api/core"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	logger *zap.Logger
	api    core.API
	router *chi.Mux
}

func NewServer(logger *zap.Logger, api core.API) *Server {
	return &Server{
		logger: logger,
		api:    api,
		router: chi.NewRouter(),
	}
}

func (s *Server) Start() {

	// Set Middleware
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.AllowContentType("application/json"))
	s.router.Use(middleware.Timeout(time.Second * 60)) // need to select ctx.Done() channel to enforce deadline
	s.router.Use(middleware.CleanPath)

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	s.router.Mount("/kills", KillsRoutes())

	http.ListenAndServe(":3000", s.router)
}

/////////////////// // // / / /  Handle Kills

type KillsHandler struct{}

func (k *KillsHandler) GetKills(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "100 Kills")
}

func (k *KillsHandler) GetKill(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "10 Kills for %s", id)
}

func KillsRoutes() chi.Router {
	r := chi.NewRouter()
	killsHandler := KillsHandler{}
	r.Get("/", killsHandler.GetKills)
	r.Get("/{id}", killsHandler.GetKill)
	return r
}
