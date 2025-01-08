package database_interface

import "github.com/jariinc/dosetti/internal/data"

type TenantRepository interface {
	FindById(id int) (data.Tenant, error)
	FindByUUID(uuid string) (data.Tenant, error)
	Save(tenant *data.Tenant) error
}
