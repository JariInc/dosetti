package session

import (
	"crypto/rand"
	"math/big"

	"github.com/jariinc/dosetti/internal/data"
)

type Session struct {
	Tenant *data.Tenant
	Key    string
}

func NewSession() Session {
	key := createNewKey()

	session := Session{
		Tenant: data.NewTenant(key),
		Key:    key,
	}

	return session
}

func createNewKey() string {
	var i big.Int
	key_bytes := make([]byte, 8)

	rand.Read(key_bytes)
	i.SetBytes(key_bytes[:])

	return i.Text(62)
}
