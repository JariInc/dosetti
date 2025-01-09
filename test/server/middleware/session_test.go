package middleware_test

import (
	"testing"

	"github.com/jariinc/dosetti/internal/server/middleware"
	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	s := middleware.NewSession()
	assert.NotEqual(t, s.Key, nil)
}

func TestLoadSession(t *testing.T) {
	base62 := "7GYVBwj8jeuPNvO87zBA58"
	s := middleware.Session{Key: base62}

	assert.Equal(t, s.Key, base62)
}
