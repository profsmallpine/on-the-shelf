package web

import (
	"net/http"

	"github.com/profsmallpine/books/html"
)

func (h *Handler) getLogin(w http.ResponseWriter, r *http.Request) {
	html.UnauthenticatedLayout(
		h.flashes(w, r),
		html.Login(),
	).Render(r.Context(), w)
}
