package repositories

import (
	"context"

	"github.com/andremeirelesmaxx/go-template/internal/core/database/gen"
)

type Example struct {
	q *gen.Queries
}

func NewExample(q *gen.Queries) Example {
	return Example{q}
}

func (e Example) ByID(id int64) (gen.Example, error) {
	return e.q.GetExampleByID(context.Background(), id)
}
