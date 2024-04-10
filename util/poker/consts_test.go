package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTexasBase(t *testing.T) {
	b := GetTexasBase()
	assert.Equal(t, 52, len(b))
}
