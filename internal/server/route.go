package server

import (
	"net/http"

	"github.com/jariinc/dosetti/internal/database"
)

func AddRoutes(
	mux *http.ServeMux,
	repos *database.Repositories) {
	mux.Handle("GET /{$}", RedirectToDayView())
	mux.Handle("GET /{key}/{$}", RenderDayView())
	mux.Handle("GET /{key}/day/{day}/{$}", RenderBody(repos))
	mux.Handle("POST /{key}/day/{day}/serving", RenderServing(repos))
}
