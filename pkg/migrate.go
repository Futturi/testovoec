package pkg

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrat(host string) error {
	time.Sleep(time.Second)
	m, err := migrate.New(
		"file://migrations", fmt.Sprintf(
			"postgres://postgres:postgres@%s:5432/postgres?sslmode=disable", host))
	if err != nil {
		return err
	}
	return m.Up()
}
