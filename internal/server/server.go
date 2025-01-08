package server

import (
	"net/http"
	"time"

	"github.com/MadAppGang/httplog"
	"github.com/jariinc/dosetti/internal/database"
	"github.com/jariinc/dosetti/internal/server/middleware"
)

func NewServer(db *database.Database) http.Handler {
	mux := http.NewServeMux()
	repos := database.NewRepositories(db)

	AddRoutes(mux, repos)

	loggingMiddleware := httplog.LoggerWithFormatter(httplog.DefaultLogFormatter)
	sessionMiddleware := middleware.SessionMiddleware(repos.TenantRepository)

	var handler http.Handler = mux
	handler = loggingMiddleware(handler)
	handler = sessionMiddleware(handler)
	handler = http.TimeoutHandler(handler, 10*time.Second, "503 Service Unavailable")

	return handler
}
