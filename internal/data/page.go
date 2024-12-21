package data

import (
	"time"
)

type Page struct {
	TenantId int
	Servings []*Serving
	Date     time.Time
}

func NewPage(tenantId int, date time.Time) *Page {
	page := &Page{
		Date:     date,
		TenantId: tenantId,
		Servings: []*Serving{},
	}

	return page
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
