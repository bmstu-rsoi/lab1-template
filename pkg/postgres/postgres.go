package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Astemirdum/person-service/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	goose "github.com/pressly/goose/v3"
)

type DB struct {
	Host     string `yaml:"host" envconfig:"DB_HOST"`
	Port     int    `yaml:"port" envconfig:"DB_PORT"`
	Username string `yaml:"user" envconfig:"DB_USER"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	NameDB   string `yaml:"dbname" envconfig:"DB_NAME"`
}

func NewPostgresDB(cfg *DB) (*sqlx.DB, error) {
	if err := migrateSchema(cfg); err != nil {
		return nil, err
	}
	return newDB(cfg)
}

func newDB(cfg *DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", newDSN(cfg))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db master ping: %w", err)
	}
	return db, nil
}

func newDSN(cfg *DB) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password)
}

func migrateSchema(cfg *DB) error {
	dsn := newDSN(cfg)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return fmt.Errorf("migrateSchema ping: %w", err)
	}

	goose.SetBaseFS(migrations.MigrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err = goose.Up(db, "sql"); err != nil {
		return errors.Wrap(err, "goose run()")
	}
	return nil
}
