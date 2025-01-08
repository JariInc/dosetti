package libsql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database/database_interface"
)

type LibSQLPrescriptionRepository struct {
	Database *Connection
}

func NewLibSQLPrescriptionRepository(db *Connection) database_interface.PrescriptionRepository {
	return &LibSQLPrescriptionRepository{Database: db}
}

func (repo *LibSQLPrescriptionRepository) FindById(tenantId int, id int) (*data.Prescription, error) {
	var prescription data.Prescription
	var start_date_str string
	var end_date_str sql.NullString
	var err error

	row := repo.Database.Conn.QueryRow("SELECT * FROM prescription WHERE tenant = ? AND id = ?", tenantId, id)
	if err = row.Scan(
		&prescription.Id,
		&prescription.TenantId,
		&prescription.Interval,
		&prescription.IntervalUnit,
		&start_date_str,
		&end_date_str,
		&prescription.Medicine,
		&prescription.MedicineAmount,
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
	rows, err := repo.Database.Conn.Query(`
		SELECT
			id,
		    tenant,
			interval,
			interval_unit,
			start_at,
			end_at,
			medicine,
			amount
		FROM prescription
		WHERE
			tenant = ?
			AND (
				(start_at BETWEEN ? AND ?)
				OR (end_at BETWEEN ? AND ?)
			 	OR (start_at <= ? AND end_at > ?)
				OR (start_at <= ? AND end_at IS NULL)
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
