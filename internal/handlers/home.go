package handler

import (
	"net/http"

	"forum/internal/cookies"
	"forum/internal/render"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, r, http.StatusNotFound)

		return
	}
	_, errC := cookies.GetCookie(r)

	var data bool
	if errC != nil {
		data = false
	} else {
		data = true
	}
	switch r.Method {
	case http.MethodGet:
		render.Render(w, "ui/html/index.html", render.WebPage{
			IsLoggedin: data,
		})

	default:
		ErrorPage(w, r, http.StatusMethodNotAllowed)
	}
}
