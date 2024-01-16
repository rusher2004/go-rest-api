package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DataStore interface {
	DeleteUser(context.Context, DeleteUserInput) error
	GetUser(context.Context, GetUserInput) (User, error)
	GetUserList(context.Context, GetUserListInput) ([]User, error)
	PostUser(context.Context, PostUserInput) (User, error)
	PutUser(context.Context, PutUserInput) (User, error)
}

type Server struct {
	newStore DataStore
	oldStore DataStore

	router http.Handler
}

type HTTPError struct {
	status  int
	message string
}

func (e HTTPError) Error() string { return e.message }

var ErrNotImplemented = HTTPError{http.StatusNotImplemented, "Endpoint not implemented by processor"}

func respond(w http.ResponseWriter, r *http.Request, data any, err error, status int) {
	type response struct {
		Data  any    `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json")
	out := response{}

	if err != nil {
		var hErr HTTPError
		if errors.As(err, &hErr) {
			w.WriteHeader(hErr.status)
			out.Error = hErr.message
		} else {
			w.WriteHeader(status)
			out.Error = "unknown error"
		}
	}

	w.WriteHeader(status)
	out.Data = data

	_ = json.NewEncoder(w).Encode(out)
}

func (s *Server) Router() http.Handler {
	return s.router
}

func (s *Server) routes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(s.withProcParam)

	r.Route("/user", func(userRoute chi.Router) {
		userRoute.Get("/", s.handleGetUserList)
		userRoute.Post("/", s.handlePostUser)

		userRoute.Route("/{id}", func(userIDRoute chi.Router) {
			userIDRoute.Use(withUserID)
			userIDRoute.Delete("/", s.handleDeleteUser)
			userIDRoute.Get("/", s.handleGetUser)
			userIDRoute.Put("/", s.handlePutUser)
		})
	})

	s.router = r
}

func NewServer(new, old DataStore) Server {
	s := Server{
		newStore: new,
		oldStore: old,
	}

	s.routes()

	return s
}
