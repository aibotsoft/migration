package migration

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"time"
)

// Up migration
func Up() error {
	start := time.Now()
	dbUrl, err := MustGetEnv("POSTGRES_URL")
	if err != nil {
		return err
	}
	soursUrl, err := MustGetEnv("MIGRATE_SOURCE_URL")
	if err != nil {
		return err
	}

	m, err := migrate.New(soursUrl, dbUrl)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	elapsed := time.Since(start)
	log.Print(elapsed)
	return nil
}
