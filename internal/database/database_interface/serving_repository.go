package database_interface

import "github.com/jariinc/dosetti/internal/data"

type ServingRepository interface {
	FindByOccurrence(tenantId int, prescriptionId int, occurrence int) (*data.Serving, error)
	FindByOccurrences(tenantId int, prescriptionId int, occurrence []int) ([]*data.Serving, error)
	Save(serving *data.Serving) error
}
