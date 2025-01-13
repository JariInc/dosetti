package server

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jariinc/dosetti/internal/database"
	"github.com/jariinc/dosetti/internal/page"
	"github.com/jariinc/dosetti/internal/server/middleware"
	assets "github.com/jariinc/dosetti/web"
)

var tmpl, _ = template.ParseGlob("web/html/*.html")

func RedirectToDayView() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			session := ctx.Value("session").(middleware.Session)
			current_day := time.Now()
			url := fmt.Sprintf("/%s/%s/", session.Key, current_day.Format("2006-01-02"))
			http.Redirect(w, r, url, http.StatusSeeOther)
		},
	)
}

func RenderDayView() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value("session").(middleware.Session)
			date, err := time.Parse("2006-01-02", r.PathValue("date"))

			if err != nil {
				date = time.Now()
			}

			err = tmpl.ExecuteTemplate(w, "index.html", struct {
				CSS        template.CSS
				JavaScript template.JS
				SessionKey string
				CurrentDay time.Time
			}{
				CSS:        template.CSS(assets.CSS),
				JavaScript: template.JS(assets.JavaScript),
				SessionKey: session.Key,
				CurrentDay: date,
			})

			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)
}

func RenderBody(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value("session").(middleware.Session)
			date, err := time.Parse("2006-01-02", r.PathValue("date"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			page := page.NewPage(repos, session, date)

			if err := tmpl.ExecuteTemplate(w, "body.html", page); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)
}

func RenderServing(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tenantId, err := strconv.Atoi(r.URL.Query().Get("tenant"))
			if err != nil {
				http.Error(w, "Unable to parse tenant", http.StatusBadRequest)
				return
			}

			prescriptionId, err := strconv.Atoi(r.URL.Query().Get("prescription"))
			if err != nil {
				http.Error(w, "Unable to parse prescription", http.StatusBadRequest)
				return
			}

			occurrence, err := strconv.Atoi(r.URL.Query().Get("occurrence"))
			if err != nil {
				http.Error(w, "Unable to parse occurrence", http.StatusBadRequest)
				return
			}

			var taken bool

			switch r.URL.Query().Get("taken") {
			case "true":
				taken = true
			case "false":
				taken = false
			default:
				http.Error(w, "Invalid taken value", http.StatusBadRequest)
			}

			serving, err := repos.ServingRepository.FindByOccurrence(1, prescriptionId, occurrence)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				} else {
					prescription, err := repos.PresciptionRepostiory.FindById(tenantId, prescriptionId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusNotFound)
						return
					}

					serving = prescription.NewServing(occurrence)
				}
			}

			serving.Taken = taken
			serving.TakenAt = time.Now()

			err = repos.ServingRepository.Save(serving)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "serving.html", serving)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}
