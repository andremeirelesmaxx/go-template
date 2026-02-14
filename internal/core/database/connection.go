package database

import (
	"database/sql"
	"fmt"

	"github.com/andremeirelesmaxx/go-template/internal/core/env"
	_ "github.com/go-sql-driver/mysql"
)

func NewConn(env env.Env) *sql.DB {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DB.User,
		env.DB.Password,
		env.DB.Host,
		env.DB.Port,
		env.DB.Name,
	)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic("could not connect to database")
	}

	return db
}
