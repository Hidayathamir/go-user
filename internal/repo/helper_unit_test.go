package repo

import (
	"time"

	"github.com/Masterminds/squirrel"
)

// contains helper for unit test.

// https://github.com/pashagolub/pgxmock?tab=readme-ov-file#matching-arguments-like-timetime
type anyTime struct{}

// Match satisfies sqlmock.Argument interface.
func (a anyTime) Match(v interface{}) bool {
	_, ok := v.(time.Time)
	return ok
}

var builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
