package server

import (
	"net/http"

	"github.com/jariinc/dosetti/internal/database"
)

func AddRoutes(
	mux *http.ServeMux,
	repos *database.Repositories) {
	mux.Handle("GET /", RenderFrontpage())
	mux.Handle("GET /partials/body", RenderBody(repos))
}
