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
	"github.com/jariinc/dosetti/internal/server/session"
	assets "github.com/jariinc/dosetti/web"
)

var tmpl, _ = template.ParseGlob("web/html/*.html")

func RedirectToDayView() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value("session").(session.Session)
			url := fmt.Sprintf("/%s/%s/", session.Key, time.Now().Format("2006-01-02"))
			http.Redirect(w, r, url, http.StatusSeeOther)
		},
	)
}

func RenderDayView(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value("session").(session.Session)
			date, err := time.Parse("2006-01-02", r.PathValue("date"))

			if err != nil {
				date = time.Now()
			}

			if r.Header.Get("HX-Request") == "true" {
				page := page.NewPage(repos, session, date)

				if err := tmpl.ExecuteTemplate(w, "body.html", page); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
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
			}
		},
	)
}

func RenderServing(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value("session").(session.Session)

			date, err := time.Parse("2006-01-02", r.PathValue("date"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			prescriptionId, err := strconv.Atoi(r.PathValue("prescription"))
			if err != nil {
				http.Error(w, "Unable to parse prescription", http.StatusBadRequest)
				return
			}

			occurrence, err := strconv.Atoi(r.PathValue("occurrence"))
			if err != nil {
				http.Error(w, "Unable to parse occurrence", http.StatusBadRequest)
				return
			}

			amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
			if err != nil {
				http.Error(w, "Unable to parse amount", http.StatusBadRequest)
				return
			}

			var taken bool
			var doses_used float64 = 0.0

			switch r.PathValue("taken") {
			case "taken":
				taken = true
			case "not-taken":
				taken = false
			default:
				http.Error(w, "Invalid taken value", http.StatusBadRequest)
			}

			serving, err := repos.ServingRepository.FindByOccurrence(session.Tenant.Id, prescriptionId, occurrence)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				} else {
					prescription, err := repos.PresciptionRepostiory.FindById(session.Tenant.Id, prescriptionId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusNotFound)
						return
					}

					serving = prescription.NewServing(occurrence)

					if taken {
						doses_used = +amount
					}
				}
			} else {
				switch taken {
				case true:
					doses_used = +amount
				case false:
					doses_used = -amount
				}

			}

			serving.Taken = taken
			serving.TakenAt = time.Now()
			err = repos.ServingRepository.Save(serving)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if doses_used != 0 {
				medicine, err := repos.MedicineRepostory.FindById(session.Tenant.Id, serving.Medicine)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				medicine.DosesLeft -= doses_used

				err = repos.MedicineRepostory.Save(medicine)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			page := page.NewPage(repos, session, date)
			err = tmpl.ExecuteTemplate(w, "serving.html", page)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}
