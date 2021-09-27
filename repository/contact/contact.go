package contact

import (
	"database/sql"

	"github.com/mohammaderm/PhoneBook/models"
)

type (
	MysqlRepo struct {
		db *sql.DB
	}
	MainRepository interface {
		Create() error
		Delete() error
		GetByName(name string) (*models.Cantact, error)
		GetAll() ([]*models.Cantact, error)
	}
)

var (
	CreateContactTable = "CREATE TABLE IF NOT EXISTS contacts(id int NOT NULL AUTO_INCREMENT, name varchar(20), lastname varchar(20), address varchar(100), phonenumber varchar(10), categoryid int, PRIMARY KEY (id), FOREIGN KEY (categoryid) REFERENCES category(categoryid), ON UPDATE CASCADE ON DELETE CASCADE)"
	CreateCategoryTable = "CREATE TABLE IF NOT EXISTS category(id int NOT NULL AUTO_INCREMENT, name varchar(20), PRIMARY KEY (id))" 
)



func NewMysqlRepo(db *MysqlRepo) (MainRepository, error) {
	repo := &MysqlRepo{}
	if err := repo.connect(); err != nil {
		return nil, err
	}
	_, err := db.db.Exec(CreateCategoryTable)
	if err != nil{
		return nil, err
	}

	_, err = db.db.Exec(CreateContactTable)
	if err != nil{
		return nil, err
	}
	return repo, nil

}