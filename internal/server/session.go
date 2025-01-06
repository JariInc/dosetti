package server

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

type Session struct {
	TenantId int
	UUID     uuid.UUID
}

func NewSession() *Session {
	s := Session{}
	var err error
	s.UUID, err = uuid.NewRandom()

	if err != nil {
		panic("cannot create uuid")
	}

	return &s
}

func LoadSession(base62_uuid string) (*Session, error) {
	var i big.Int
	_, ok := i.SetString(base62_uuid, 62)
	if !ok {
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
