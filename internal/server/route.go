package server

import (
	"net/http"

	"github.com/jariinc/dosetti/internal/database"
)

func AddRoutes(
	mux *http.ServeMux,
	repos *database.Repositories) {
	mux.Handle("GET /{$}", RedirectToDayView())
	mux.Handle("GET /{key}/{date}/{$}", RenderDayView())
	mux.Handle("GET /{key}/{date}/servings/{$}", RenderBody(repos))
	mux.Handle("POST /{key}/{date}/servings/{servingId}", RenderServing(repos))
}
