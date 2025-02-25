package libsql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database/database_interface"
)

type LibSQLServingRepository struct {
	Database *sql.DB
}

func NewLibSQLServingRepository(db *sql.DB) database_interface.ServingRepository {
	return &LibSQLServingRepository{Database: db}
}

func (repo *LibSQLServingRepository) FindByOccurrence(tenantId int, prescriptionId int, occurrence int) (*data.Serving, error) {
	var serving data.Serving
	var taken_at_str sql.NullString

	row := repo.Database.QueryRow(`
		SELECT s.id, s.tenant, s.prescription, s.occurrence, s.amount, s.taken, s.taken_at, p.medicine, m.name, m.doses_left
		FROM serving AS s
		JOIN prescription AS p ON s.prescription = p.id
		JOIN medicine AS m ON p.medicine = m.id
		WHERE s.tenant = ? AND s.prescription = ? AND s.occurrence = ?`,
		tenantId, prescriptionId, occurrence)

	if err := row.Scan(
		&serving.Id,
		&serving.TenantId,
		&serving.PrescriptionId,
		&serving.Occurrence,
		&serving.MedicineAmount,
		&serving.Taken,
		&taken_at_str,
		&serving.Medicine,
		&serving.MedicineName,
		&serving.DosesLeft,
	); err != nil {
		if err == sql.ErrNoRows {
			return &data.Serving{}, fmt.Errorf("ServingByOccurrence %d %d %d: %w", tenantId, prescriptionId, occurrence, err)
		}

		if taken_at_str.Valid {
			serving.TakenAt, err = time.Parse(DATE_TIME_FORMAT, taken_at_str.String)

			if err != nil {
				return &data.Serving{}, fmt.Errorf("ServingByOccurrence %d %d %d: %w", tenantId, prescriptionId, occurrence, err)
			}
		}

		return &data.Serving{}, fmt.Errorf("ServingByOccurrence %d %d %d: %w", tenantId, prescriptionId, occurrence, err)
	}
	return &serving, nil

}

func (repo *LibSQLServingRepository) FindByOccurrences(tenantId int, prescriptionId int, occurrence []int) ([]*data.Serving, error) {
	var servings []*data.Serving
	var occurrences_str []string
	var occurrences_sql string

	// TODO: Refactor to generic function
	for _, occurance := range occurrence {
		occurrences_str = append(occurrences_str, fmt.Sprintf("%d", occurance))
	}

	occurrences_sql = strings.Join(occurrences_str, ", ")

	query := fmt.Sprintf(`
		SELECT s.id, s.tenant, s.prescription, s.occurrence, s.taken, s.taken_at, p.medicine, s.amount, m.name, m.doses_left
		FROM serving AS s
		JOIN prescription AS p ON s.prescription = p.id
		JOIN medicine AS m ON p.medicine = m.id
		WHERE s.tenant = ? AND s.prescription = ? AND s.occurrence IN (%s)`, occurrences_sql)

	rows, err := repo.Database.Query(query, tenantId, prescriptionId)

	if err != nil {
		return []*data.Serving{}, fmt.Errorf("ServingsByOccurences %d %d: %v", tenantId, prescriptionId, err)
	}

	defer rows.Close()

	for rows.Next() {
		var serving data.Serving
		var taken_at_str sql.NullString

		if err := rows.Scan(
			&serving.Id,
			&serving.TenantId,
			&serving.PrescriptionId,
			&serving.Occurrence,
			&serving.Taken,
			&taken_at_str,
			&serving.Medicine,
			&serving.MedicineAmount,
			&serving.MedicineName,
			&serving.DosesLeft,
		); err != nil {
			return []*data.Serving{}, err
		}

		if taken_at_str.Valid {
			serving.TakenAt, err = time.Parse(DATE_TIME_FORMAT, taken_at_str.String)
			if err != nil {
				return []*data.Serving{}, fmt.Errorf("ServingsByOccurences %d %d: %v", tenantId, prescriptionId, err)
			}
		}

		servings = append(servings, &serving)
	}

	return servings, nil
}

func (repo *LibSQLServingRepository) Save(serving *data.Serving) error {
	var taken_at sql.NullString

	if !serving.TakenAt.IsZero() {
		taken_at = sql.NullString{
			String: serving.TakenAt.Format(DATE_TIME_FORMAT),
			Valid:  true,
		}
	}

	result, err := repo.Database.Exec(`
		REPLACE INTO serving
			(tenant, prescription, occurrence, amount, taken, taken_at)
			VALUES (?, ?, ?, ?, ?, ?)`,
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
