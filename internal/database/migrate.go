package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	sqlitemigrate "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func Open(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)

	const pragmas = `PRAGMA foreign_keys = ON;
	PRAGMA journal_mode = WAL;
	PRAGMA synchronous = NORMAL;
	PRAGMA locking_mode = NORMAL;
	PRAGMA busy_timeout = 10000;`
	_, err = db.Exec(pragmas)
	if err != nil {
		return nil, fmt.Errorf("pragmas: %w", err)
	}

	return &DB{db: db}, nil
}

//go:embed migrations/*.sql
var migrationsFS embed.FS

func (db *DB) Migrate() error {
	driver, err := sqlitemigrate.WithInstance(db.db, &sqlitemigrate.Config{})
	if err != nil {
		return fmt.Errorf("get sqlite instance: %w", err)
	}

	source, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("create iofs source: %w", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs", source,
		"sqlite", driver,
	)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}

func (db *DB) Close() error {
	return db.db.Close()
}
