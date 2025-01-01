package server

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jariinc/dosetti/internal/database"
	"github.com/jariinc/dosetti/internal/page"
	assets "github.com/jariinc/dosetti/web"
)

var tmpl = template.Must(template.ParseGlob("./web/html/*.html"))

func RenderFrontpage() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.ExecuteTemplate(w, "index.html", struct {
				CSS        template.CSS
				JavaScript template.JS
			}{
				CSS:        template.CSS(assets.CSS),
				JavaScript: template.JS(assets.JavaScript),
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}

func RenderBody(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//date := time.Now()
			tenantId, err := strconv.Atoi(r.URL.Query().Get("t"))

			if err != nil {
				tenantId = 1
			}

			page := page.NewPage(repos, tenantId, time.Now())

			// prescriptions, err := repos.PresciptionRepostiory.FindByTenant(tenantId)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }

			// var servings []*data.Serving

			// for _, prescription := range prescriptions {
			// 	servings = append(servings, prescription.NewServing(date))
			// }

			// page.Servings = servings

			err = tmpl.ExecuteTemplate(w, "body.html", page)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}

func RenderServing(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			servingId, err := strconv.Atoi(r.URL.Query().Get("id"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			serving, err := repos.ServingRepository.FindById(1, servingId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			err = tmpl.ExecuteTemplate(w, "serving.html", serving)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}
