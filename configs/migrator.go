package configs

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func ExecMigrations(db *sqlx.DB) error {
	maxVersion := 1
	err := db.Get(&maxVersion, "SELECT MAX(version) FROM migrations")
	if err != nil {
		return err
	}

	migrations, err := findMigrationsToRun(maxVersion)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		sqlx.LoadFile(db, "./migrations/"+migration)
	}

	return nil
}

func findMigrationsToRun(maxVersion int) ([]string, error) {
	migrationsDir := "./migrations/"
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	var migrations []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			version, err := strconv.Atoi(file.Name()[:1])
			if err != nil {
				return nil, err
			}

			if version > maxVersion {
				migrations = append(migrations, file.Name())
			}
		}
	}

	return migrations, nil
}
