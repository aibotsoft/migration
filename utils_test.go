package migration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustGetEnv_equal(t *testing.T) {
	got, err := MustGetEnv("TEST_ENV_VAR")
	assert.Equal(t, got, "OK")
	assert.Nil(t, err)
}
func TestMustGetEnv_not_exists(t *testing.T) {
	got, err := MustGetEnv("NOT_EXISTENT_ENV")
	assert.Equal(t, got, "")
	assert.Error(t, err)
}
