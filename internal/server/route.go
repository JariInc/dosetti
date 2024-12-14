package server

import (
	"net/http"

	"github.com/jariinc/dosetti/internal/database"
)

func AddRoutes(
	mux *http.ServeMux,
	tenant_repo *database.TenantRepository,
	prescription_repo *database.PrescriptionRepository) {
	mux.Handle("GET /", RenderFrontpage())
	mux.Handle("GET /partials/body", RenderBody(tenant_repo))
}
