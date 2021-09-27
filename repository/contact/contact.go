package contact

import (
	"database/sql"
)

type (
	MysqlRepo struct {
		db *sql.DB
	}
	MainRepository interface {
		Create(name, lastname, address, phonenumber string, categoryid int) error
		// Delete() error
		// GetByName(name string) (*models.Cantact, error)
		// GetAll() ([]*models.Cantact, error)
	}
)

var (
	CreateContactTable  = "CREATE TABLE IF NOT EXISTS contact(id int NOT NULL AUTO_INCREMENT, name varchar(20), lastname varchar(20), address varchar(100), phonenumber varchar(10), categoryid int, PRIMARY KEY (id), FOREIGN KEY (categoryid) REFERENCES category(categoryid), ON UPDATE CASCADE ON DELETE CASCADE)"
	CreateCategoryTable = "CREATE TABLE IF NOT EXISTS category(id int NOT NULL AUTO_INCREMENT, name varchar(20), PRIMARY KEY (id))"
	InsertContact       = "INSERT INTO contact (name, lastname, address, phonenumber, categoryid) VALUES (?,?,?,?,?)"
)

func NewMysqlRepo() (MainRepository, error) {
	repo := &MysqlRepo{}
	if err := repo.connect(); err != nil {
		return nil, err
	}
	_, err := repo.db.Exec(CreateCategoryTable)
	if err != nil {
		return nil, err
	}

	_, err = repo.db.Exec(CreateContactTable)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (db *MysqlRepo) Create(name, lastname, address, phonenumber string, categoryid int) error {
	_, err := db.db.Exec(InsertContact, name, lastname, address, phonenumber, categoryid)
	if err != nil {
		return err
	}
	return nil
}
