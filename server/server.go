package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/core"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	logger *zap.Logger
	config *config.Config
	api    core.API
	router *chi.Mux
}

func NewServer(logger *zap.Logger, conf *config.Config, api core.API) *Server {
	return &Server{
		logger: logger,
		config: conf,
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

	s.router.Mount("/kills", KillsRoutes(s.api))

	s.logger.Info("Starting server on port " + s.config.Options.Service.HttpPort)
	http.ListenAndServe(":"+s.config.Options.Service.HttpPort, s.router)
}

/////////////////// // // / / /  Handle Kills

type KillsHandler struct {
	api core.API
}

func (k *KillsHandler) GetKills(w http.ResponseWriter, r *http.Request) {
	res, err := k.api.GetAllKills()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	total := 0
	for _, v := range res {
		total += v
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.Itoa(total))
}

func (k *KillsHandler) GetKill(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	total, err := k.api.GetKills(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strconv.Itoa(total))
}

func KillsRoutes(api core.API) chi.Router {
	r := chi.NewRouter()
	killsHandler := KillsHandler{api: api}
	r.Get("/", killsHandler.GetKills)
	r.Get("/{id}", killsHandler.GetKill)
	return r
}
