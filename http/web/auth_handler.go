package web

import (
	"net/http"

	"github.com/xy-planning-network/trails/http/resp"
)

func (h *Handler) getAuth_GoogleLogin(w http.ResponseWriter, r *http.Request) {
	rt, err := h.Auth.FetchAuthURL(w)
	if err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}
	w.Header().Add("HX-Redirect", rt)
}

func (h *Handler) getAuth_GoogleCallback(w http.ResponseWriter, r *http.Request) {
	data, err := h.Auth.ExchangeCode(r)
	if err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	user, err := h.User.HandleCallback(data)
	if err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	s, err := h.session(r.Context())
	if err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	if err := s.RegisterUser(w, r, user.ID); err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Code(http.StatusSeeOther), resp.Url("/"))
		return
	}

	h.Redirect(w, r, resp.Success("Welcome, nice to see you here!"), resp.Url(user.HomePath()))
}
