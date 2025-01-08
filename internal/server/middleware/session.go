package middleware

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/jariinc/dosetti/internal/data"
)

type Session struct {
	Tenant *data.Tenant
	UUID   uuid.UUID
}

func NewSession() *Session {
	tenant_uuid, err := uuid.NewRandom()

	if err != nil {
		panic("cannot create uuid")
	}

	s := Session{
		Tenant: data.NewTenant(tenant_uuid.String()),
		UUID:   tenant_uuid,
	}

	return &s
}

func LoadSession(base62_uuid string) (*Session, error) {
	var i big.Int
	_, err := i.SetString(base62_uuid, 62)
	if !err {
		return &Session{}, fmt.Errorf("cannot parse base62: %q", base62_uuid)
	}

	var uuid uuid.UUID
	copy(uuid[:], i.Bytes())

	s := Session{
		UUID: uuid,
	}

	return &s, nil
}

func (s *Session) Base62UUID() string {
	var i big.Int
	i.SetBytes(s.UUID[:])
	return i.Text(62)
}
