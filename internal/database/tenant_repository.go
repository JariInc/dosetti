package database

import (
	"database/sql"
	"fmt"

	"github.com/jariinc/dosetti/internal/data"
)

type TenantRepository struct {
	Database *Database
}

func NewTenantRepository(db *Database) *TenantRepository {
	return &TenantRepository{Database: db}
}

func (repo *TenantRepository) FindById(id int) (data.Tenant, error) {
	var tenant data.Tenant
	row := repo.Database.Conn.QueryRow("SELECT * FROM tenant WHERE id = ?", id)
	if err := row.Scan(&tenant.Id, &tenant.UUID); err != nil {
		if err == sql.ErrNoRows {
			return tenant, fmt.Errorf("TenantById %d: no such tenant", id)
		}
		return tenant, fmt.Errorf("TenantById %d: %v", id, err)
	}
	return tenant, nil
}

func (repo *TenantRepository) FindByUUID(uuid string) (data.Tenant, error) {
	var tenant data.Tenant
	row := repo.Database.Conn.QueryRow("SELECT * FROM tenant WHERE uuid = ?", uuid)
	if err := row.Scan(&tenant.Id, &tenant.UUID); err != nil {
		if err == sql.ErrNoRows {
			return tenant, fmt.Errorf("TenantByUUID %s: no such tenant", uuid)
		}
		return tenant, fmt.Errorf("TenantByUUID %s: %v", uuid, err)
	}
	return tenant, nil
}

func (repo *TenantRepository) Save(tenant *data.Tenant) error {
	result, err := repo.Database.Conn.Exec(`
		REPLACE INTO tenant
			(uuid)
			VALUES (?)`,
		tenant.UUID,
	)

	if err != nil {
		return fmt.Errorf("Save serving: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Save serving: %v", err)
	}
	tenant.Id = int(id)

	return nil
}
