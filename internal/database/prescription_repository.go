package database

import (
	"database/sql"
	"fmt"

	"github.com/jariinc/dosetti/internal/data"
)

type PrescriptionRepository struct {
	Database *Database
}

func NewPrescriptionRepository(db *Database) PrescriptionRepository {
	return PrescriptionRepository{Database: db}
}

func (repo *PrescriptionRepository) FindById(id int) (data.Prescription, error) {
	var prescription data.Prescription
	row := repo.Database.Conn.QueryRow("SELECT * FROM prescription WHERE id = ?", id)
	if err := row.Scan(
		&prescription.Id,
		&prescription.Tenant.Id,
		&prescription.Interval,
		&prescription.IntervalUnit,
		&prescription.StartDate,
		&prescription.Offset,
		&prescription.Medicine,
		&prescription.MedicineAmount,
		&prescription.MedicineUnit,
	); err != nil {
		if err == sql.ErrNoRows {
			return prescription, fmt.Errorf("PrescriptionById %d: no such prescription", id)
		}
		return prescription, fmt.Errorf("PrescriptionById %d: %v", id, err)
	}
	return prescription, nil
}

func (repo *PrescriptionRepository) FindByTenant(tenantId int) (data.Prescription, error) {
	var prescription data.Prescription
	row := repo.Database.Conn.QueryRow("SELECT * FROM prescription WHERE tenant = ?", tenantId)
	if err := row.Scan(
		&prescription.Id,
		&prescription.Tenant.Id,
		&prescription.Interval,
		&prescription.IntervalUnit,
		&prescription.StartDate,
		&prescription.Offset,
		&prescription.Medicine,
		&prescription.MedicineAmount,
		&prescription.MedicineUnit,
	); err != nil {
		if err == sql.ErrNoRows {
			return prescription, fmt.Errorf("PrescriptionByTenant %d: no such prescription", tenantId)
		}
		return prescription, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
	}
	return prescription, nil
}
