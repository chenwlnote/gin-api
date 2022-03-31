package dao

import "database/sql"

type MysqlWriter interface {
	GetWriter() *sql.DB
}
