package models

type (
	// for post contact detail in post request
	ReqCreateContact struct {
		Name        string `json:"name" db:"name"`
		LastName    string `json:"lastname" db:"lastname"`
		Address     string `json:"address" db:"address"`
		PhoneNumber string `json:"phonenumber" db:"phonenumber"`
		CategoryID  uint   `json:"categoryid" db:"categoryid"`
	}

	// for Interaction with database
	Cantact struct {
		Id          uint   `json:"id" db:"id"`
		Name        string `json:"name" db:"name"`
		LastName    string `json:"lastname" db:"lastname"`
		Address     string `json:"address" db:"address"`
		PhoneNumber string `json:"phonenumber" db:"phonenumber"`
		CategoryID  uint   `json:"categoryid" db:"categoryid"`
		UserID      uint   `json:"userid" db:"userid"`
	}

	// for get contact detail(category name , id) in get request
	ReqGetContact struct {
		Id           uint   `json:"id" db:"id"`
		Name         string `json:"name" db:"name"`
		LastName     string `json:"lastname" db:"lastname"`
		Address      string `json:"address" db:"address"`
		PhoneNumber  string `json:"phonenumber" db:"phonenumber"`
		UserID       uint   `json:"userid" db:"userid"`
		CategoryID   uint   `json:"categoryid" db:"categoryid"`
		CategoryName string `json:"categoryname" db:"categoryname"`
	}
)
