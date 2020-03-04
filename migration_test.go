package migration

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("Tests begin")
	_ = os.Setenv("TEST_ENV_VAR", "OK")
	exitVal := m.Run()
	_ = os.Unsetenv("TEST_ENV_VAR")
	os.Exit(exitVal)
}

func TestUp_no_dbUrl(t *testing.T) {
	expectedError := errors.New("POSTGRES_URL env is not set")
	err := Up()
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}
func TestUp_no_soursUrl(t *testing.T) {
	_ = os.Setenv("POSTGRES_URL", "SET")
	defer os.Unsetenv("POSTGRES_URL")
	expectedError := errors.New("MIGRATE_SOURCE_URL env is not set")
	err := Up()
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}
