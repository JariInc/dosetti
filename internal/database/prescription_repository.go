package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jariinc/dosetti/internal/data"
)

type PrescriptionRepository struct {
	Database *Database
}

func NewPrescriptionRepository(db *Database) *PrescriptionRepository {
	return &PrescriptionRepository{Database: db}
}

func (repo *PrescriptionRepository) FindById(id int) (data.Prescription, error) {
	var prescription data.Prescription
	row := repo.Database.Conn.QueryRow("SELECT * FROM prescription WHERE id = ?", id)
	if err := row.Scan(
		&prescription.Id,
		&prescription.TenantId,
		&prescription.Interval,
		&prescription.IntervalUnit,
		&prescription.StartDate,
		&prescription.Offset,
		&prescription.Medicine,
		&prescription.MedicineAmount,
	); err != nil {
		if err == sql.ErrNoRows {
			return prescription, fmt.Errorf("PrescriptionById %d: no such prescription", id)
		}
		return prescription, fmt.Errorf("PrescriptionById %d: %v", id, err)
	}
	return prescription, nil
}

func (repo *PrescriptionRepository) FindByTenant(tenantId int) ([]data.Prescription, error) {
	var err error
	var prescriptions []data.Prescription
	rows, err := repo.Database.Conn.Query("SELECT id, tenant, interval, interval_unit, start_date, offset, medicine, amount FROM prescription WHERE tenant = ?", tenantId)

	if err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
	}

	defer rows.Close()

	for rows.Next() {
		var prescription data.Prescription
		var start_date_str string
		if err := rows.Scan(
			&prescription.Id,
			&prescription.TenantId,
			&prescription.Interval,
			&prescription.IntervalUnit,
			&start_date_str,
			&prescription.Offset,
			&prescription.Medicine,
			&prescription.MedicineAmount,
		); err != nil {
			return []data.Prescription{}, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
		}
		prescription.StartDate, err = time.Parse("2006-01-02", start_date_str)

		prescriptions = append(prescriptions, prescription)
	}

	if err := rows.Err(); err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
	}

	if err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
	}

	return prescriptions, nil
}
