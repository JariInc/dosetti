package database_interface

import (
	"time"

	"github.com/jariinc/dosetti/internal/data"
)

type PrescriptionRepository interface {
	FindById(tenantId int, id int) (*data.Prescription, error)
	FindBetweenDates(tenantId int, from time.Time, to time.Time) ([]data.Prescription, error)
}
