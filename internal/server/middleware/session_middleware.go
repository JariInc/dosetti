package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

const TENANT_COOKIE = "tenant"

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session *Session
		cookie, err := r.Cookie(TENANT_COOKIE)

		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				session = NewSession()
			default:
				log.Println(err)
				http.Error(w, "cookie error", http.StatusInternalServerError)
				return
			}
		} else {
			session, err = LoadSession(cookie.Value)

			if err != nil {
				log.Println(err)
				http.Error(w, "session loading error", http.StatusInternalServerError)
				return
			}
		}

		req := r.WithContext(context.WithValue(r.Context(), "session", session))
		*r = *req

		newCookie := http.Cookie{
			Name:     TENANT_COOKIE,
			Value:    session.Base62UUID(),
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24 * 365),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}

		next.ServeHTTP(w, r)
		http.SetCookie(w, &newCookie)

	})
}
