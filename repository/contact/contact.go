package contact

import (
	"database/sql"
	"github.com/mohammaderm/PhoneBook/models"
)


const (
	username = "mohammad"
	password = "1115444123Mohammad@"
	hostname = "127.0.0.1:3306"
	dbname = "phonebook"
)

type (
	MysqlRepo struct {
		db sql.DB
	}
	MainRepository interface{
		Create() error
		Delete() error
		GetByName(name string) (*models.Cantact, error)
		GetAll()([]*models.Cantact, error)
	}
)
 

func NewMysqlRepo(db *MysqlRepo) (MainRepository , error){
	
}