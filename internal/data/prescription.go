package data

import (
	"database/sql/driver"
	"time"
)

type IntervalUnit string

func (u *IntervalUnit) Scan(value interface{}) error { *u = IntervalUnit(value.(string)); return nil }
func (u IntervalUnit) Value() (driver.Value, error)  { return string(u), nil }

const (
	IntervalHourly  = "hourly"
	IntervalDaily   = "daily"
	IntervalWeekly  = "weekly"
	IntervalMonthly = "monthly"
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
