package database_interface

import (
	"github.com/jariinc/dosetti/internal/data"
)

type MedicineRepository interface {
	FindById(tenantId int, id int) (*data.Medicine, error)
	Save(medicine *data.Medicine) error
}
