package data

import "time"

type IntervalUnit int

const (
	Hourly IntervalUnit = iota
	Daily
	Weekly
	Monthly
)

type Prescription struct {
	Id             int
	TenantId       int
	Interval       int
	IntervalUnit   IntervalUnit
	StartDate      time.Time
	Offset         int
	Medicine       string
	MedicineAmount string
}

func (p Prescription) NewServing(date time.Time) *Serving {
	serving := &Serving{
		Prescription: p,
		ScheduledAt:  date,
		Taken:        false,
	}

	return serving
}
