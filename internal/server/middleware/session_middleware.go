package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jariinc/dosetti/internal/database/database_interface"
)

const TENANT_COOKIE = "tenant"

func SessionMiddleware(tenant_repo database_interface.TenantRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var session Session
			cookie, err := r.Cookie(TENANT_COOKIE)

			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					session = NewSession()
					tenant_repo.Save(session.Tenant)
				default:
					log.Println(err)
					http.Error(w, "cookie error", http.StatusInternalServerError)
					return
				}
			} else {
				session = Session{Key: cookie.Value}

				tenant, err := tenant_repo.FindByKey(session.Key)
				if err != nil {
					session = NewSession()
					tenant_repo.Save(session.Tenant)
				} else {
					session.Tenant = &tenant
				}
			}

			req := r.WithContext(context.WithValue(r.Context(), "session", session))
			*r = *req

			newCookie := http.Cookie{
				Name:     TENANT_COOKIE,
				Value:    session.Key,
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
