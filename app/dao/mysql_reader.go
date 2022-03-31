package dao

import "database/sql"

type MysqlReader interface {
	GetReader() *sql.DB
}
