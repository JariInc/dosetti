package page

import (
	"log"
	"time"

	"github.com/jariinc/dosetti/internal/database"
)

func NewPage(repos *database.Repositories, tenantId int, date time.Time) *Page {
	page := &Page{
		Date:     date,
		TenantId: tenantId,
	}

	from := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	to := from.Add(time.Hour * 24)

	prescriptions, err := repos.PresciptionRepostiory.FindBetweenDates(tenantId, from, to)
	if err != nil {
		log.Fatal(err)
	}

	for _, prescription := range prescriptions {
		occurrances := prescription.OccurrancesBetweenDates(from, to)
		servings, servings_err := repos.ServingRepository.FindByOccurrences(tenantId, prescription.Id, occurrances)

		if servings_err != nil {
			log.Fatal(servings_err)
		}

		for _, occurrance := range occurrances {
			found := false
			for _, serving := range servings {
				if serving.Occurrence == occurrance {
					page.Servings = append(page.Servings, serving)
					found = true
				}
			}

			if !found {
				page.Servings = append(page.Servings, prescription.NewServing(occurrance))

			}
		}
	}

	return page
}
