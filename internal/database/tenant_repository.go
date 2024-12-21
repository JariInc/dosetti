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
	if err := row.Scan(&tenant.Id, &tenant.Key); err != nil {
		if err == sql.ErrNoRows {
			return tenant, fmt.Errorf("TenantById %d: no such tenant", id)
		}
		return tenant, fmt.Errorf("TenantById %d: %v", id, err)
	}
	return tenant, nil
}

func (repo *TenantRepository) FindByKey(key string) (data.Tenant, error) {
	var tenant data.Tenant
	row := repo.Database.Conn.QueryRow("SELECT * FROM tenant WHERE key = ?", key)
	if err := row.Scan(&tenant.Id, &tenant.Key); err != nil {
		if err == sql.ErrNoRows {
			return tenant, fmt.Errorf("TenantByKey %s: no such tenant", key)
		}
		return tenant, fmt.Errorf("TenantByKey %s: %v", key, err)
	}
	return tenant, nil
}
