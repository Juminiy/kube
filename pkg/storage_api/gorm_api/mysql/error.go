package mysql

import (
	mysqldriver "github.com/go-sql-driver/mysql"
)

func FullGroupBy(err error) bool {
	return errIs(err, 1055, [5]byte{'4', '2', '0', '0', '0'}) ||
		errIs(err, 1140, [5]byte{'4', '2', '0', '0', '0'})
}

func errIs(err error, errNumber uint16, sqlState [5]byte) bool {
	if errv, ok := err.(*mysqldriver.MySQLError); ok {
		return errv.Number == errNumber &&
			errv.SQLState == sqlState
	}
	return false
}
