package web

import (
	"net/http"

	"github.com/xy-planning-network/trails/http/router"
)

func (h *Handler) RegisterRoutes() {
	unauthenticatedRoutes := []router.Route{
		{Path: "/auth/google/callback", Method: http.MethodGet, Handler: h.getAuth_GoogleCallback},
		{Path: "/auth/google/login", Method: http.MethodGet, Handler: h.getAuth_GoogleLogin},
		{Path: "/", Method: http.MethodGet, Handler: h.getLogin},
	}
	h.UnauthedRoutes(unauthenticatedRoutes)

	authenticatedRoutes := []router.Route{
		{Path: "/books", Method: http.MethodGet, Handler: h.getBooks},
		{Path: "/logoff", Method: http.MethodGet, Handler: h.getLogoff},
	}
	h.AuthedRoutes("/", "/logoff", authenticatedRoutes)
}
