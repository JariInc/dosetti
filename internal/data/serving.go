package data

import "time"

type Serving struct {
	Id             int
	TenantId       int
	PrescriptionId int
	Occurrence     int
	Medicine       int
	MedicineName   string
	MedicineAmount float64
	Taken          bool
	TakenAt        time.Time
}
