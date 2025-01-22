package data

type Tenant struct {
	Id  int
	Key string
}

func NewTenant(key string) *Tenant {
	return &Tenant{
		Key: key,
	}
}
