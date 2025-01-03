package data

import (
	"database/sql/driver"
	"math"
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
	StartAt        time.Time
	EndAt          time.Time
	Medicine       string
	MedicineAmount string
}

func (p *Prescription) NewServing(occurrence int) *Serving {
	return &Serving{
		TenantId:       p.TenantId,
		PrescriptionId: p.Id,
		Occurrence:     occurrence,
		MedicineAmount: p.MedicineAmount,
		Taken:          false,
	}
}

func (p *Prescription) OccurrancesBetweenDates(from time.Time, to time.Time) []int {
	hours_to_from := from.Sub(p.StartAt).Hours()   // inclusive start
	hours_to_to := (to.Sub(p.StartAt) - 1).Hours() // exclusive end
	var occurrence_from int
	var occurrence_to int
	var occurences []int

	switch p.IntervalUnit {
	case IntervalHourly:
		occurrence_from = int(math.Ceil(hours_to_from / float64(p.Interval)))
		occurrence_to = int(math.Floor(hours_to_to / float64(p.Interval)))
	case IntervalDaily:
		occurrence_from = int(math.Ceil(hours_to_from / (float64(p.Interval) * 24)))
		occurrence_to = int(math.Floor(hours_to_to / (float64(p.Interval) * 24)))
	case IntervalWeekly:
		occurrence_from = int(math.Ceil(hours_to_from / (float64(p.Interval) * 24 * 7)))
		occurrence_to = int(math.Floor(hours_to_to / (float64(p.Interval) * 24 * 7)))
	case IntervalMonthly:
		//nextMonth := iter.AddDate(0, 1, 0)
		//interval = nextMonth.Sub(iter)
		// TODO
	default:
		panic("unknown interval unit")
	}

	for i := occurrence_from; i <= occurrence_to; i++ {
		occurences = append(occurences, i)
	}

	return occurences
}
