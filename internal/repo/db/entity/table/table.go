// Package table should be use to get table name and column name when query to database.
package table

// init initiate all table name and column name. It's oke to use init here
// because it's just initialize table and column name.
func init() { //nolint:gochecknoinits
	initTableUser()
}
