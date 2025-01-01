package data

import "time"

type Serving struct {
	Id             int
	TenantId       int
	PrescriptionId int
	Occurrence     int
	MedicineAmount string
	Taken          bool
	TakenAt        time.Time
}
