package web

import (
	"net/http"

	"github.com/profsmallpine/on-the-shelf/html"
	"github.com/xy-planning-network/trails/http/resp"
)

func (h *Handler) getLogin(w http.ResponseWriter, r *http.Request) {
	html.UnauthenticatedLayout(
		h.flashes(w, r),
		html.Login(),
	).Render(r.Context(), w)
}

func (h *Handler) getLogoff(w http.ResponseWriter, r *http.Request) {
	s, err := h.session(r.Context())
	if err != nil {
		h.Redirect(w, r, resp.Err(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	if err := s.Delete(w, r); err != nil {
		h.Redirect(w, r, resp.Err(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	w.Header().Add("HX-Redirect", "/")
}
