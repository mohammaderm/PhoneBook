package models

type (
	Cantact struct {
		Name       string `json:"name" db:"name"`
		LastName   string `json:"lastname" db:"lastname"`
		Address    string `json:"address" db:"address"`
		CategoryID uint   `json:"categoryid" db:"categoryid"`
	}
	Category struct {
		Name string `json:"name" db:"name"`
	}
)
