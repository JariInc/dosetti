package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database"
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
				log.Fatal(err)
			}
		},
	)
}

func RenderBody(repos *database.Repositories) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			date := time.Now()
			tenantId, err := strconv.Atoi(r.URL.Query().Get("t"))

			if err != nil {
				tenantId = 1
			}

			page := data.NewPage(tenantId, time.Now())

			prescriptions, err := repos.PresciptionRepostiory.FindByTenant(tenantId)
			if err != nil {
				log.Fatal(err)
			}

			var servings []*data.Serving

			for _, prescription := range prescriptions {
				servings = append(servings, prescription.NewServing(date))
			}

			page.Servings = servings

			err = tmpl.ExecuteTemplate(w, "body.html", page)

			if err != nil {
				log.Fatal(err)
			}
		},
	)
}
