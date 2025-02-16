package web

import (
	"net/http"

	"github.com/xy-planning-network/trails/http/router"
)

func (h *Handler) Router() {
	unauthenticatedRoutes := []router.Route{
		{Path: "/", Method: http.MethodGet, Handler: h.getLogin},
	}
	h.UnauthedRoutes(unauthenticatedRoutes)
}
