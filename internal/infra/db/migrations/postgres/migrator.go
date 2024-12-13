package postgres

import (
	"context"
	"database/sql"
	effective_mobile_tz "effective-mobile-tz"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

type DefaultMigrator struct {
	db      *sql.DB
	adapter *ZlgLogAdapter
}

func NewDefaultMigrator(db *sql.DB, adapter *ZlgLogAdapter) *DefaultMigrator {
	return &DefaultMigrator{db: db, adapter: adapter}
}

func (d *DefaultMigrator) Migrate(ctx context.Context) error {
	goose.SetBaseFS(effective_mobile_tz.Migrations)
	goose.SetLogger(d.adapter)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose set dialect: %w", err)
	}

	if err := goose.Up(d.db, "migrations"); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}

type ZlgLogAdapter struct {
	logger *zerolog.Logger
}

func NewZlgLogAdapter(logger *zerolog.Logger) *ZlgLogAdapter {
	return &ZlgLogAdapter{logger: logger}
}

func (z *ZlgLogAdapter) Fatalf(format string, v ...interface{}) {
	z.logger.Fatal().Msgf(format, v...)
}

func (z *ZlgLogAdapter) Printf(format string, v ...interface{}) {
	z.logger.Error().Msgf(format, v...)
}
