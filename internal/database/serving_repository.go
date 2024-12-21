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

func (repo *ServingRepository) FindByDate(tenantId int, date time.Time) (data.Serving, error) {
	var serving data.Serving
	row := repo.Database.Conn.QueryRow("SELECT * FROM serving WHERE tenant = ? AND DATE(scheduled_at) = ?", tenantId, date.Format("2006-01-02"))
	if err := row.Scan(
		&serving.Id,
		&serving.TenantId,
		&serving.Interval,
		&serving.IntervalUnit,
		&serving.StartDate,
		&serving.Offset,
		&serving.Medicine,
		&serving.MedicineAmount,
		&serving.Taken,
		&serving.TakenAt,
		&serving.ScheduledAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return serving, fmt.Errorf("PrescriptionByTenant %d: no such prescription", tenantId)
		}
		return serving, fmt.Errorf("PrescriptionByTenant %d: %v", tenantId, err)
	}
	return serving, nil
}
