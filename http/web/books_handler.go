package web

import (
	"net/http"

	"github.com/profsmallpine/on-the-shelf/html"
	"github.com/xy-planning-network/trails/http/resp"
)

func (h *Handler) getBooks(w http.ResponseWriter, r *http.Request) {
	user, err := h.currentUser(r.Context())
	if err != nil {
		h.Redirect(w, r, resp.Url("/logoff"))
		return
	}

	html.AuthenticatedLayout(
		h.flashes(w, r),
		html.ShowBooks(),
		user,
	).Render(r.Context(), w)
}
