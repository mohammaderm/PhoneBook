package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "mohammad"
	password = "1115444123Mohammad@"
	hostname = "127.0.0.1:3306"
	dbname   = "phonebook"
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
