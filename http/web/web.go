package web

import (
	"context"
	"net/http"

	"github.com/profsmallpine/books/domain"
	"github.com/xy-planning-network/trails"
	"github.com/xy-planning-network/trails/http/session"
	"github.com/xy-planning-network/trails/ranger"
)

type Handler struct {
	*ranger.Ranger
}

// session helps by retrieving the session.TrailsSessionable from the provided context.
func (h *Handler) session(ctx context.Context) (session.Session, error) {
	s, ok := ctx.Value(trails.SessionKey).(session.Session)
	if !ok {
		return session.Session{}, domain.ErrNoSession
	}
	return s, nil
}

// flashes helps by retrieving the session.Flashes from the provided context.
func (h *Handler) flashes(w http.ResponseWriter, r *http.Request) []session.Flash {
	s, _ := h.session(r.Context())
	return s.Flashes(w, r)
}
