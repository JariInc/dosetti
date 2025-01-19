package server

import (
	"net/http"

	"github.com/jariinc/dosetti/internal/database"
)

func AddRoutes(
	mux *http.ServeMux,
	repos *database.Repositories) {
	mux.Handle("GET /{$}", RedirectToDayView())
	mux.Handle("GET /{key}/{$}", RedirectToDayView())
	mux.Handle("GET /{key}/{date}/{$}", RenderDayView(repos))
	mux.Handle("POST /{key}/{date}/servings/prescription/{prescription}/occurrence/{occurrence}/{taken}", RenderServing(repos))
}
