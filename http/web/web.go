package web

import (
	"context"
	"net/http"

	"github.com/profsmallpine/on-the-shelf/db/models"
	"github.com/profsmallpine/on-the-shelf/domain"
	"github.com/xy-planning-network/trails"
	"github.com/xy-planning-network/trails/http/session"
	"github.com/xy-planning-network/trails/ranger"
)

type Handler struct {
	domain.Procedures
	*models.Queries
	*ranger.Ranger
	domain.Services
}

// currentUser helps by retrieving the User stored from the provided context.
func (h *Handler) currentUser(ctx context.Context) (*domain.User, error) {
	u, ok := ctx.Value(trails.CurrentUserKey).(*domain.User)
	if !ok {
		return nil, domain.ErrNoUser
	}
	return u, nil
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
