package infrastructure

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

type PostgresDatabase struct {
	db *sqlx.DB
}

func NewPostgresDb(databaseURL string) (*PostgresDatabase, error) {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{db}, nil
}

func (pg PostgresDatabase) RunMigrations() error {
	driver, err := postgres.WithInstance(pg.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db_migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
