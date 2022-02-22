package repository

import (
	"database/sql"

	"github.com/mohammaderm/PhoneBook/models"
)

type (
	MysqlRepo struct {
		db *sql.DB
	}
	MainRepository interface {

		// user interface
		Create_User(use *models.User) error
		Delete_User(id string) (bool, error)
		GetbyEmail_User(email string) (*models.User, error)
		GetByUserName_User(username string) (*models.User, error)
		// ------------------------------------------------------------

		// contacts interface
		Create_contact(cont *models.Cantact) error
		Delete_contact(id, userid string) (bool, error)
		GetByName_contact(name, userid string) ([]models.ReqGetContact, error)
		GetAll_contact(userid string) ([]models.ReqGetContact, error)
		GetByCategoryId_contact(id, userid string) ([]models.ReqGetContact, error)
		// ------------------------------------------------------------

		// category interface
		Create_category(cate *models.Category) error
		Delete_category(id, userid string) (bool, error)
		GetAll_category(userid string) ([]models.Category, error)
		// ------------------------------------------------------------
	}
)

var (
	// contacts querys
	InsertContact          = "INSERT INTO contact (name, lastname, address, phonenumber, categoryid, userid) VALUES (?,?,?,?,?,?)"
	DeleteContact          = "DELETE FROM contact WHERE id=? AND userid=?"
	GetContactByName       = "SELECT contact.*, category.name FROM contact INNER JOIN category ON contact.categoryid = category.id WHERE contact.name = ? OR contact.lastname = ? AND userid = ?"
	GetAllContact          = "SELECT contact.*, category.name FROM contact INNER JOIN category ON contact.categoryid = category.id WHERE contact.userid=?"
	GetContactByCategoryId = "SELECT contact.*, category.name FROM contact INNER JOIN category ON contact.categoryid = category.id WHERE categoryid=? AND contact.userid=?"
	// ------------------------------------------------------------

	// categorys querys
	InsertCategory = "INSERT INTO category (name, userid) VALUES (?,?)"
	DeleteCategory = "DELETE FROM category WHERE id=? AND userid=?"
	GetAllCategory = "SELECT * FROM category WHERE userid=?"
	// ------------------------------------------------------------

	// user querys
	CreateUser        = "INSERT INTO user (username, email, password) VALUES (?,?,?)"
	DeleteUser        = "DELETE FROM user WHERE id=?"
	GetUserbyEmail    = "SELECT * FROM user WHERE email=?"
	GetUserbyusername = "SELECT * FROM user WHERE username=?"
	// ------------------------------------------------------------
)

func NewMysqlRepo() (MainRepository, error) {
	repo := &MysqlRepo{}
	if err := repo.connect(); err != nil {
		return nil, err
	}
	return repo, nil
}

// --------------------------
// all User repositorys
// --------------------------
func (db *MysqlRepo) Create_User(use *models.User) error {
	_, err := db.db.Exec(CreateUser, use.Username, use.Email, use.Password)
	if err != nil {
		return err
	}
	return nil
}

func (db *MysqlRepo) Delete_User(id string) (bool, error) {
	result, err := db.db.Exec(DeleteUser, id)
	if err != nil {
		return false, err
	}
	cont, _ := result.RowsAffected()
	if cont == 0 {
		return false, err
	}
	return true, nil
}

func (db *MysqlRepo) GetbyEmail_User(email string) (*models.User, error) {
	var result models.User
	err := db.db.QueryRow(GetUserbyEmail, email).Scan(&result.Username, &result.Email, &result.Password, &result.Id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *MysqlRepo) GetByUserName_User(username string) (*models.User, error) {
	var result models.User
	err := db.db.QueryRow(GetUserbyEmail, username).Scan(&result.Username, &result.Email, &result.Password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// --------------------------
// all Contact repositorys
// --------------------------
func (db *MysqlRepo) Create_contact(cont *models.Cantact) error {
	_, err := db.db.Exec(InsertContact, cont.Name, cont.LastName, cont.Address, cont.PhoneNumber, cont.CategoryID, cont.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *MysqlRepo) Delete_contact(id, userid string) (bool, error) {
	result, err := db.db.Exec(DeleteContact, id, userid)
	if err != nil {
		return false, err
	}
	cont, _ := result.RowsAffected()
	if cont == 0 {
		return false, nil
	}
	return true, nil
}

func (db *MysqlRepo) GetByName_contact(name, userid string) ([]models.ReqGetContact, error) {
	var result []models.ReqGetContact
	rows, err := db.db.Query(GetContactByName, name, name, userid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res models.ReqGetContact
		if err := rows.Scan(&res.Id,
			&res.Name,
			&res.LastName,
			&res.Address,
			&res.PhoneNumber,
			&res.CategoryID,
			&res.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

func (db *MysqlRepo) GetAll_contact(userid string) ([]models.ReqGetContact, error) {
	var result []models.ReqGetContact
	rows, err := db.db.Query(GetAllContact, userid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res models.ReqGetContact
		if err := rows.Scan(
			&res.Name,
			&res.LastName,
			&res.Address,
			&res.Id,
			&res.CategoryID,
			&res.UserID,
			&res.PhoneNumber,
			&res.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

func (db *MysqlRepo) GetByCategoryId_contact(id, userid string) ([]models.ReqGetContact, error) {
	var result []models.ReqGetContact
	rows, err := db.db.Query(GetContactByCategoryId, id, userid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res models.ReqGetContact
		if err := rows.Scan(
			&res.Name,
			&res.LastName,
			&res.Address,
			&res.Id,
			&res.CategoryID,
			&res.UserID,
			&res.PhoneNumber,
			&res.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

// --------------------------
// all category repository
// --------------------------
func (db *MysqlRepo) Create_category(cate *models.Category) error {
	_, err := db.db.Exec(InsertCategory, cate.Name, cate.UserID)
	if err != nil {
		return err
	}
	return nil
}
func (db *MysqlRepo) Delete_category(id, userid string) (bool, error) {
	result, err := db.db.Exec(DeleteCategory, id, userid)
	if err != nil {
		return false, err
	}
	cont, _ := result.RowsAffected()
	if cont == 0 {
		return false, err
	}
	return true, nil
}
func (db *MysqlRepo) GetAll_category(userid string) ([]models.Category, error) {
	var result []models.Category
	rows, err := db.db.Query(GetAllCategory, userid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res models.Category
		if err := rows.Scan(&res.Name, &res.Id, &res.UserID); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}
