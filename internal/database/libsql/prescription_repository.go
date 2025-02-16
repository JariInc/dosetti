package libsql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database/database_interface"
)

type LibSQLPrescriptionRepository struct {
	Database *sql.DB
}

func NewLibSQLPrescriptionRepository(db *sql.DB) database_interface.PrescriptionRepository {
	return &LibSQLPrescriptionRepository{Database: db}
}

func (repo *LibSQLPrescriptionRepository) FindById(tenantId int, id int) (*data.Prescription, error) {
	var prescription data.Prescription
	var start_date_str string
	var end_date_str sql.NullString
	var err error

	row := repo.Database.QueryRow(`
		SELECT
			p.id,
			p.tenant,
		 	p.interval,
			p.interval_unit,
			p.start_at,
			p.end_at,
			p.medicine,
			p.amount,
			m.name,
			m.doses_left
		FROM prescription AS p
		JOIN medicine AS m ON p.medicine = m.id
		WHERE p.tenant = ? AND p.id = ?`, tenantId, id)
	if err = row.Scan(
		&prescription.Id,
		&prescription.TenantId,
		&prescription.Interval,
		&prescription.IntervalUnit,
		&start_date_str,
		&end_date_str,
		&prescription.Medicine,
		&prescription.MedicineAmount,
		&prescription.MedicineName,
		&prescription.DosesLeft,
	); err != nil {
		if err == sql.ErrNoRows {
			return &data.Prescription{}, fmt.Errorf("PrescriptionById %d: %v", id, err)
		}

		return &prescription, fmt.Errorf("PrescriptionById %d: %v", id, err)
	}

	prescription.StartAt, err = time.Parse(DATE_TIME_FORMAT, start_date_str)
	if err != nil {
		return &data.Prescription{}, fmt.Errorf("PrescriptionById %d: %v", id, err)
	}

	if end_date_str.Valid {
		prescription.EndAt, err = time.Parse(DATE_TIME_FORMAT, end_date_str.String)
		if err != nil {
			return &data.Prescription{}, fmt.Errorf("PrescriptionById %d: %v", id, err)
		}
	}

	return &prescription, nil
}

func (repo *LibSQLPrescriptionRepository) FindBetweenDates(tenantId int, from time.Time, to time.Time) ([]data.Prescription, error) {
	var err error
	var prescriptions []data.Prescription
	rows, err := repo.Database.Query(`
		SELECT
			p.id,
			p.tenant,
			p.interval,
			p.interval_unit,
			p.start_at,
			p.end_at,
			p.medicine,
			p.amount,
			m.name,
			m.doses_left
		FROM prescription AS p
		JOIN medicine AS m ON p.medicine = m.id
		WHERE
			p.tenant = ?
			AND (
				(p.start_at BETWEEN ? AND ?)
				OR (p.end_at BETWEEN ? AND ?)
			 	OR (p.start_at <= ? AND p.end_at > ?)
				OR (p.start_at <= ? AND p.end_at IS NULL)
			)`, tenantId, from, to, from, to, from, to, to)

	if err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
	}

	defer rows.Close()

	for rows.Next() {
		var prescription data.Prescription
		var start_date_str string
		var end_date_str sql.NullString
		if err := rows.Scan(
			&prescription.Id,
			&prescription.TenantId,
			&prescription.Interval,
			&prescription.IntervalUnit,
			&start_date_str,
			&end_date_str,
			&prescription.Medicine,
			&prescription.MedicineAmount,
			&prescription.MedicineName,
			&prescription.DosesLeft,
		); err != nil {
			return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
		}
		prescription.StartAt, err = time.Parse(DATE_TIME_FORMAT, start_date_str)
		if err != nil {
			return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
		}

		if end_date_str.Valid {
			prescription.EndAt, err = time.Parse(DATE_TIME_FORMAT, end_date_str.String)
			if err != nil {
				return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
			}
		}

		prescriptions = append(prescriptions, prescription)
	}

	if err := rows.Err(); err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
	}

	if err != nil {
		return []data.Prescription{}, fmt.Errorf("PrescriptionBetweenDates %d %s %s: %v", tenantId, from, to, err)
	}

	return prescriptions, nil
}
