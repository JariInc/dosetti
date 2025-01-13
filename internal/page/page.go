package page

import (
	"time"

	"github.com/jariinc/dosetti/internal/data"
)

type Page struct {
	TenantId    int
	Servings    []*data.Serving
	Today       time.Time
	CurrentDay  time.Time
	NextDay     time.Time
	PreviousDay time.Time
	SessionKey  string
}
