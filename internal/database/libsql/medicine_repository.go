package libsql

import (
	"database/sql"
	"fmt"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database/database_interface"
)

type LibSQLMedicineRepository struct {
	Database *sql.DB
}

func NewLibSQLMedicineRepository(db *sql.DB) database_interface.MedicineRepository {
	return &LibSQLMedicineRepository{Database: db}
}

func (repo *LibSQLMedicineRepository) FindById(tenantId int, id int) (*data.Medicine, error) {
	var medicine data.Medicine
	row := repo.Database.QueryRow(`
		SELECT id, tenant, name, doses_left
		FROM medicine
		WHERE id = ? AND tenant = ?`,
		id,
		tenantId,
	)
	if err := row.Scan(&medicine.Id, &medicine.TenantId, &medicine.Name, &medicine.DosesLeft); err != nil {
		if err == sql.ErrNoRows {
			return &medicine, fmt.Errorf("MedicinetById %d: no such medicine", id)
		}
		return &medicine, fmt.Errorf("MedicinetById %d: %v", id, err)
	}
	return &medicine, nil
}

func (repo *LibSQLMedicineRepository) Save(medicine *data.Medicine) error {
	result, err := repo.Database.Exec(`
		REPLACE INTO medicine
			(id, tenant, name, doses_left)
			VALUES (?, ?, ?, ?)`,
		medicine.Id,
		medicine.TenantId,
		medicine.Name,
		medicine.DosesLeft,
	)

	if err != nil {
		return fmt.Errorf("Save medicine: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Save medicine: %v", err)
	}
	medicine.Id = int(id)

	return nil
}
