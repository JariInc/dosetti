package data

type Tenant struct {
	Id   int
	UUID string
}

func NewTenant(uuid string) *Tenant {
	return &Tenant{
		UUID: uuid,
	}
}
