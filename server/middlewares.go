package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type processor string

const (
	newProcessor = processor("new")
	oldProcessor = processor("old")
)

// contextKey for key/vals held in request context
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "context key value: " + k.name
}

var procIfaceCTXKey = &contextKey{"procIface"}
var userIDCTXKey = &contextKey{"userID"}

// ContextValue is a shortcut to fetch the value of type T from context.
func ContextValue[T any](ctx context.Context, key *contextKey) *T {
	val, _ := ctx.Value(key).(*T)
	return val
}

// withProcParam is a middleware that ensures the proc_id query parameter exists and is valid. If valid, it will be
// inserted into the request context with context key `procIfaceCTXKey`
func (s *Server) withProcParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proc := processor(r.URL.Query().Get("proc"))

		if proc != newProcessor && proc != oldProcessor {
			err := HTTPError{http.StatusUnprocessableEntity, "proc param invalid"}
			respond(w, r, nil, err, err.status)
			return
		}

		var procIface *DataStore
		switch proc {
		case newProcessor:
			procIface = &s.newStore
		case oldProcessor:
			procIface = &s.oldStore
		}

		ctx := context.WithValue(r.Context(), procIfaceCTXKey, procIface)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func withUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		u, err := uuid.Parse(id)
		if err != nil {
			hErr := HTTPError{http.StatusUnprocessableEntity, "invalid id format"}
			respond(w, r, nil, hErr, hErr.status)
			return
		}

		ctx := context.WithValue(r.Context(), userIDCTXKey, &u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
