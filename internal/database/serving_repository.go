package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jariinc/dosetti/internal/data"
)

type ServingRepository struct {
	Database *Database
}

func NewServingRepository(db *Database) *ServingRepository {
	return &ServingRepository{Database: db}
}

func (repo *ServingRepository) FindById(tenantId int, servingId int) (*data.Serving, error) {
	var serving data.Serving
	var taken_at_str string

	row := repo.Database.Conn.QueryRow("SELECT * FROM serving WHERE tenant = ? AND id = ?", tenantId, servingId)
	if err := row.Scan(
		&serving.Id,
		&serving.TenantId,
		&serving.PrescriptionId,
		&serving.Occurrence,
		&serving.MedicineAmount,
		&serving.Taken,
		&taken_at_str,
	); err != nil {
		serving.TakenAt, err = time.Parse(DATE_TIME_FORMAT, taken_at_str)
		if err != nil {
			return &data.Serving{}, fmt.Errorf("ServingsByDate %d: %v", tenantId, err)
		}

		if err == sql.ErrNoRows {
			return &data.Serving{}, fmt.Errorf("ServingById %d: no such serving", servingId)
		}
		return &data.Serving{}, fmt.Errorf("ServingById %d: %v", servingId, err)
	}
	return &serving, nil
}

func (repo *ServingRepository) FindBetweenDates(tenantId int, prescriptionId int, from time.Time, to time.Time) ([]*data.Serving, error) {
	var servings []*data.Serving

	rows, err := repo.Database.Conn.Query("SELECT * FROM serving WHERE tenant = ? AND prescription = ? AND ", tenantId, prescriptionId, from.Format(DATE_TIME_FORMAT), to.Format(DATE_TIME_FORMAT))

	if err != nil {
		return []*data.Serving{}, fmt.Errorf("ServingsByDate %d: %v", tenantId, err)
	}

	defer rows.Close()

	for rows.Next() {
		var serving data.Serving
		var taken_at_str string

		if err := rows.Scan(
			&serving.Id,
			&serving.TenantId,
			&serving.PrescriptionId,
			&serving.Occurrence,
			&serving.MedicineAmount,
			&serving.Taken,
			&taken_at_str,
		); err != nil {
			return []*data.Serving{}, fmt.Errorf("ServingsByDate %d: %v", tenantId, err)
		}

		serving.TakenAt, err = time.Parse(DATE_TIME_FORMAT, taken_at_str)
		if err != nil {
			return []*data.Serving{}, fmt.Errorf("ServingsByDate %d: %v", tenantId, err)
		}

		servings = append(servings, &serving)
	}

	return servings, nil
}

func (repo *ServingRepository) Save(serving *data.Serving) error {
	var taken_at sql.NullString

	if serving.TakenAt.IsZero() == false {
		taken_at = sql.NullString{
			String: serving.TakenAt.Format(DATE_TIME_FORMAT),
			Valid:  true,
		}
	}

	result, err := repo.Database.Conn.Exec(`
		INSERT INTO serving
			(tenant, prescription, occurrence, amount, taken, taken_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)
			RETURNING id`,
		serving.TenantId,
		serving.PrescriptionId,
		serving.Occurrence,
		serving.MedicineAmount,
		serving.Taken,
		taken_at,
	)

	if err != nil {
		return fmt.Errorf("Save serving: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Save serving: %v", err)
	}
	serving.Id = int(id)

	return nil
}
