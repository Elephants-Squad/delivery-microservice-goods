package migrate

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	psql "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Импорт драйвера для файлов
)

const (
	sslMode          = "?sslmode=disable"
	dataBase         = "postgres"
	pathToMigrations = "file:///app/internal/data/"
)

func Up(dbURL string) error {
	db, err := sql.Open(dataBase, dbURL+sslMode)
	if err != nil {
		return err
	}

	driver, err := psql.WithInstance(db, &psql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigrations,
		dataBase,
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
