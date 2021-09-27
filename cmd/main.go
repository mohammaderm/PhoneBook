package main

import "github.com/mohammaderm/PhoneBook/repository/contact"

func main() {
	_, err := contact.NewMysqlRepo()
	if err != nil {
		panic(err)
	}
}
