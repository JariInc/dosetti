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
	log.Println("NewPage found", len(prescriptions), "prescriptions")

	//page.Servings = servings

	return page
}
