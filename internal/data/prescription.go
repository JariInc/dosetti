package data

import "time"

type IntervalUnit int

const (
	Hourly IntervalUnit = iota
	Daily
	Weekly
)

type MedicineUnit string

const (
	Gram      MedicineUnit = "g"
	Milligram MedicineUnit = "mg"
	Microgram MedicineUnit = "ug"
)

type Prescription struct {
	Id             int
	Tenant         Tenant
	Interval       int
	IntervalUnit   IntervalUnit
	StartDate      time.Time
	Offset         int
	Medicine       string
	MedicineAmount float32
	MedicineUnit   MedicineUnit
}

func (p Prescription) NewServing(date time.Time) *Serving {
	serving := &Serving{
		Prescription: p,
		ScheduledAt:  date,
		Taken:        false,
	}

	return serving
}
