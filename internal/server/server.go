package server

import (
	"net/http"

	"github.com/MadAppGang/httplog"
	"github.com/jariinc/dosetti/internal/database"
)

func NewServer(db *database.Database) http.Handler {
	mux := http.NewServeMux()
	tenant_repo := database.NewTenantRepository(db)
	prescription_repo := database.NewPrescriptionRepository(db)

	AddRoutes(mux, &tenant_repo, &prescription_repo)

	loggingMiddleware := httplog.LoggerWithFormatter(httplog.DefaultLogFormatter)

	var handler http.Handler = mux
	handler = loggingMiddleware(handler)

	return handler
}
