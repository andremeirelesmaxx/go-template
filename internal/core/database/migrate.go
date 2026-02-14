package database

import (
	"database/sql"
	"fmt"

	"github.com/andremeirelesmaxx/go-template/internal/core/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigrate(db *sql.DB, env env.Env) *migrate.Migrate {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(fmt.Sprint("could not create migration instance with existing db instance ", "error=", err.Error()))
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/database/migrations", env.App.Root),
		"mysql",
		driver,
	)

	if err != nil {
		panic(fmt.Sprint("could not create a instance", err.Error()))
	}

	return m
}
