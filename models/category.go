package models

type (
	Category struct {
		Id     uint   `json:"id" db:"id"`
		Name   string `json:"name" db:"name"`
		UserID uint   `json:"userid" db:"userid"`
	}
	ReqCreateCategory struct {
		Name string `json:"name" db:"name"`
	}
)
