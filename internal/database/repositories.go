package database

type Repositories struct {
	TenantRepository      *TenantRepository
	PresciptionRepostiory *PrescriptionRepository
	ServingRepository     *ServingRepository
}

func NewRepositories(db *Database) *Repositories {
	return &Repositories{
		TenantRepository:      NewTenantRepository(db),
		PresciptionRepostiory: NewPrescriptionRepository(db),
		ServingRepository:     NewServingRepository(db),
	}
}
