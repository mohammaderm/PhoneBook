package models

type (
	User struct {
		Id       uint   `json:"id" db:"id"`
		Username string `json:"username" db:"username"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}
	Login struct {
		// Username string `json:"username" db:"username"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}
)
