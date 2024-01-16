package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type (
	DeleteUserInput struct {
		ID uuid.UUID `json:"id"`
	}

	GetUserInput struct {
		ID uuid.UUID `json:"id"`
	}

	GetUserListInput struct{}

	PostUserInput struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	PutUserInput struct {
		ID    uuid.UUID
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	User struct {
		Email string    `json:"email"`
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
	}
)

var UserNotFound = HTTPError{http.StatusNotFound, "User not found"}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[uuid.UUID](ctx, userIDCTXKey)
	procIface := *ContextValue[DataStore](ctx, procIfaceCTXKey)

	if err := procIface.DeleteUser(ctx, DeleteUserInput{id}); err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, nil, nil, http.StatusNoContent)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[uuid.UUID](ctx, userIDCTXKey)
	procIface := *ContextValue[DataStore](ctx, procIfaceCTXKey)

	out, err := procIface.GetUser(ctx, GetUserInput{id})
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handleGetUserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	procIface := *ContextValue[DataStore](ctx, procIfaceCTXKey)

	out, err := procIface.GetUserList(ctx, GetUserListInput{})
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handlePostUser(w http.ResponseWriter, r *http.Request) {
	var in PostUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	ctx := r.Context()
	procIface := *ContextValue[DataStore](ctx, procIfaceCTXKey)

	out, err := procIface.PostUser(ctx, in)
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handlePutUser(w http.ResponseWriter, r *http.Request) {
	var in PutUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	ctx := r.Context()
	in.ID = *ContextValue[uuid.UUID](ctx, userIDCTXKey)
	procIface := *ContextValue[DataStore](ctx, procIfaceCTXKey)

	out, err := procIface.PutUser(ctx, in)
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}
