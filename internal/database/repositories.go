package database

import (
	"database/sql"

	"github.com/jariinc/dosetti/internal/database/database_interface"
	"github.com/jariinc/dosetti/internal/database/libsql"
)

type Repositories struct {
	TenantRepository      database_interface.TenantRepository
	PresciptionRepostiory database_interface.PrescriptionRepository
	ServingRepository     database_interface.ServingRepository
}

func NewLibSQLRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		TenantRepository:      libsql.NewLibSQLTenantRepository(db),
		PresciptionRepostiory: libsql.NewLibSQLPrescriptionRepository(db),
		ServingRepository:     libsql.NewLibSQLServingRepository(db),
	}
}
