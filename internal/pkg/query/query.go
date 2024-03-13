// Package query contains func related to sql query string.
package query

// Returning return sql query string RETURNING column.
func Returning(column string) string {
	return "RETURNING " + column
}
