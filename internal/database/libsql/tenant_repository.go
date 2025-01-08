package libsql

import (
	"database/sql"
	"fmt"

	"github.com/jariinc/dosetti/internal/data"
	"github.com/jariinc/dosetti/internal/database/database_interface"
)

type LibSQLTenantRepository struct {
	Database *Connection
}

func NewLibSQLTenantRepository(db *Connection) database_interface.TenantRepository {
	return &LibSQLTenantRepository{Database: db}
}

func (repo *LibSQLTenantRepository) FindById(id int) (data.Tenant, error) {
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

func (repo *LibSQLTenantRepository) FindByUUID(uuid string) (data.Tenant, error) {
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

func (repo *LibSQLTenantRepository) Save(tenant *data.Tenant) error {
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
