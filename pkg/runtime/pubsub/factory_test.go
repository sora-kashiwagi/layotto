package pubsub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFactory(t *testing.T) {
	f := NewFactory("", nil)
	assert.NotNil(t, f)
}
