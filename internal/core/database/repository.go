package database

import (
	"database/sql"

	"github.com/andremeirelesmaxx/go-template/internal/core/database/gen"
)

func NewRepository(db *sql.DB) *gen.Queries {
	return gen.New(db)
}
