package migration

import (
	"errors"
	"os"
)

func MustGetEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errors.New(key + " env is not set")
	}
	return v, nil
}
