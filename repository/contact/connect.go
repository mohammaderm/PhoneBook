package contact

import (
	"database/sql"
	"fmt"
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
		return nil
	}

	defer db.Close()
	d.db = db
	return nil
}
