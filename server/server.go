package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/core"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	logger *zap.Logger
	config *config.Config
	api    core.API
	db     core.DbClient
	router *chi.Mux
}

func NewServer(logger *zap.Logger, conf *config.Config, api core.API, db core.DbClient) *Server {
	return &Server{
		logger: logger,
		config: conf,
		api:    api,
		db:     db,
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

	s.router.Post("/signup", s.handleSignUp)
	s.router.Get("/users", s.getAllUsers)
	s.router.Post("/signin", s.handleSignIn)
	s.router.Mount("/kills", KillsRoutes(s.api, s.db))

	s.logger.Info("Starting server on port " + s.config.Options.Service.HttpPort)
	http.ListenAndServe(":"+s.config.Options.Service.HttpPort, s.router)
}

func KillsRoutes(api core.API, db core.DbClient) chi.Router {
	r := chi.NewRouter()
	killsHandler := KillsHandler{api: api, db: db}
	r.Get("/", killsHandler.GetKills)
	r.Get("/{id}", killsHandler.GetKill)
	return r
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// request json:
// {
// "Username": "newuser",
// "Email":"newuser@new.clom",
// "Passkey":"password"
// }
func (s *Server) handleSignUp(w http.ResponseWriter, r *http.Request) {
	var newUserData core.User
	err := json.NewDecoder(r.Body).Decode(&newUserData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if s.api.CheckUserExists(newUserData.Email) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "email alread registered")
		return
	}
	hash, err := HashPassword(newUserData.Hash)
	newUserData.Hash = hash
	newUserData.UserId = uuid.New()
	// store in db
	err = s.db.CreateUser(newUserData)
	w.WriteHeader(http.StatusOK)
}

// request json:
// {
// "Email":"newuser@new.clom",
// "Passkey":"password"
// }
func (s *Server) handleSignIn(w http.ResponseWriter, r *http.Request) {
	type signInData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var reqData signInData
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !s.api.CheckUserExists(reqData.Email) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email not registered with api")
		return
	}

	hash, err := s.db.GetHashForEmail(reqData.Email)
	if !CheckPasswordHash(reqData.Password, hash) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid password")
	}

	// issue JWT token

	w.WriteHeader(http.StatusOK)
}

func (s *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

/////////////////// // // / / /  Handle Kills

type KillsHandler struct {
	api core.API
	db  core.DbClient
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
