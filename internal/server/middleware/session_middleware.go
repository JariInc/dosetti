package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/jariinc/dosetti/internal/database/database_interface"
	"github.com/jariinc/dosetti/internal/server/session"
)

const TENANT_COOKIE = "tenant"

func SessionMiddleware(tenant_repo database_interface.TenantRepository) func(next http.Handler) http.Handler {
	match_key, _ := regexp.Compile("^/([a-zA-Z0-9]+)/")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ses session.Session

			// Match tenant key manually since ServeMux has not populated r.PathValue yet
			if match_key.MatchString(r.URL.Path) {
				key := match_key.FindStringSubmatch(r.URL.Path)[1]
				ses = session.Session{Key: key}
				tenant, err := tenant_repo.FindByKey(ses.Key)

				if err != nil {
					ses = session.NewSession()
					tenant_repo.Save(ses.Tenant)
					http.Redirect(w, r, "/"+ses.Tenant.Key+"/", http.StatusSeeOther)
				} else {
					ses.Tenant = &tenant
				}
			} else {
				cookie, err := r.Cookie(TENANT_COOKIE)

				if err != nil {
					switch {
					case errors.Is(err, http.ErrNoCookie):
						ses = session.NewSession()
						tenant_repo.Save(ses.Tenant)
					default:
						log.Println(err)
						http.Error(w, "cookie error", http.StatusInternalServerError)
						return
					}
				} else {
					ses = session.Session{Key: cookie.Value}
					tenant, err := tenant_repo.FindByKey(ses.Key)
					if err != nil {
						ses = session.NewSession()
						tenant_repo.Save(ses.Tenant)
					} else {
						ses.Tenant = &tenant
					}
				}
			}

			req := r.WithContext(context.WithValue(r.Context(), "session", ses))
			*r = *req

			newCookie := http.Cookie{
				Name:     TENANT_COOKIE,
				Value:    ses.Key,
				Path:     "/",
				Expires:  time.Now().Add(time.Hour * 24 * 365),
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}

			next.ServeHTTP(w, r)
			http.SetCookie(w, &newCookie)
		})
	}
}
