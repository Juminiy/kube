package postgres

import "github.com/jackc/pgx/v5/pgconn"

func GroupingError(err error) bool {
	return errIs(err, "42803")
}

func UndefinedColumn(err error) bool {
	return errIs(err, "42703")
}

func UniqueViolation(err error) bool {
	return errIs(err, "23505")
}

func errIs(err error, sqlState string) bool {
	if errV, ok := err.(*pgconn.PgError); ok {
		return errV.SQLState() == sqlState
	}
	return false
}
