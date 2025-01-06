package session_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jariinc/dosetti/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	s := server.NewSession()
	assert.NotEqual(t, s.UUID, uuid.Nil)
}

func TestLoadSession(t *testing.T) {
	base62 := "7GYVBwj8jeuPNvO87zBA58"
	s, err := server.LoadSession(base62)

	assert.Nil(t, err)
	assert.Equal(t, s.UUID.String(), "fcc6d2e2-5906-4f1a-a360-bac76958c2f6")
}
