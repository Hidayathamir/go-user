package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"github.com/Hidayathamir/go-user/config"
	_ "github.com/jackc/pgx/v5/stdlib" // don't really understand, remove if you know what you do, i just following this article about pgx to sql.DB. https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql#hello-world-from-postgresql
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

// MigrateUp migrate database using internal/usecase/repo/db/schema_migration.
func MigrateUp() error {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.PG.Username, config.PG.Password, config.PG.Host, config.PG.Port, config.PG.DBName,
	)

	db, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logrus.Warnf("sql.DB.Close: %v", err)
		}
	}()

	migrate.SetTable("migrations")

	var countMigrationApplied int
	for i := 0; i < 10; i++ {
		countMigrationApplied, err = migrate.Exec(db, "postgres", getFileMigrationSource(), migrate.Up)
		if err == nil {
			break
		}

		logrus.
			WithField("attempt count", i+1).
			Warnf("err migrate: migrate.Exec: %v", err)

		time.Sleep(time.Second)
	}

	if err != nil {
		return fmt.Errorf("err 10 times when try to migrate: %w", err)
	}

	logrus.Infof("migrate done, %d migration applied 🟢", countMigrationApplied)

	return nil
}

func getFileMigrationSource() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join("internal", "usecase", "repo", "db", "schema_migration"),
	}
	return migrations
}
