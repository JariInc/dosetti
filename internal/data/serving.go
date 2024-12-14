package data

import "time"

type Serving struct {
	Prescription
	Taken       bool
	TakenAt     time.Time
	ScheduledAt time.Time
}
