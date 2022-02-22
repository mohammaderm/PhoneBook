package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "#"
	password = "#"
	hostname = "#"
	dbname   = "#"
)

func (d *MysqlRepo) connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		hostname,
		dbname,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	d.db = db
	return nil
}
