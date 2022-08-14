package main

import "github.com/mohammaderm/PhoneBook/application"

func main() {
	err := application.Run()
	if err != nil {
		panic(err.Error())
	}
}
