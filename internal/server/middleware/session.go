package middleware

import (
	"crypto/rand"
	"math/big"

	"github.com/jariinc/dosetti/internal/data"
)

type Session struct {
	Tenant *data.Tenant
	Key    string
}

func NewSession() *Session {
	key := createNewKey()

	s := Session{
		Tenant: data.NewTenant(key),
		Key:    key,
	}

	return &s
}

func createNewKey() string {
	var i big.Int
	key_bytes := make([]byte, 16)

	rand.Read(key_bytes)
	i.SetBytes(key_bytes[:])

	return i.Text(62)
}
