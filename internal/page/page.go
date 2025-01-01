package page

import (
	"time"

	"github.com/jariinc/dosetti/internal/data"
)

type Page struct {
	TenantId int
	Servings []*data.Serving
	Date     time.Time
}

func (p *Page) NextDayFormatted(format string) string {
	next := p.Date.Add(time.Hour * 24)
	return next.Format(format)
}

func (p *Page) PreviousDayFormatted(format string) string {
	next := p.Date.Add(time.Hour * -24)
	return next.Format(format)
}

func (p *Page) TodayFormatted(format string) string {
	today := time.Now()
	return today.Format(format)
}
