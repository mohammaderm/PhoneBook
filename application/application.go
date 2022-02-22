package application

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammaderm/PhoneBook/handler"
	"github.com/mohammaderm/PhoneBook/repository"
	"github.com/rs/cors"
)

func Run() error {
	repo, err := repository.NewMysqlRepo()
	if err != nil {
		return err
	}
	handlerfuctions := handler.NewHandler(repo)

	r := mux.NewRouter()
	user_route := r.PathPrefix("/api/v1/").Subrouter()
	user_route.Use(handler.Auth)

	// Contact Routes
	user_route.HandleFunc("/contact/create", handlerfuctions.Create_contact).Methods("POST")
	user_route.HandleFunc("/contact/delete/{id}", handlerfuctions.Delete_contact).Methods("DELETE")
	user_route.HandleFunc("/contact/getbyname", handlerfuctions.GetByName_contact).Methods("GET")
	user_route.HandleFunc("/contact/getall", handlerfuctions.GetAll_contact).Methods("GET")
	user_route.HandleFunc("/contact/getbycategory/{id}", handlerfuctions.GetByCategory_contact).Methods("GET")

	// Category Routes
	user_route.HandleFunc("/category/create", handlerfuctions.Create_category).Methods("POST")
	user_route.HandleFunc("/category/delete/{id}", handlerfuctions.Delete_category).Methods("DELETE")
	user_route.HandleFunc("/category/getall", handlerfuctions.GetAll_category).Methods("GET")

	// Auth Routes
	r.HandleFunc("/login", handlerfuctions.Login).Methods("Post")
	r.HandleFunc("/register", handlerfuctions.Register).Methods("Post")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	fmt.Println("server is runing on port 8089...")
	http.ListenAndServe(":8089", corsWrapper.Handler(r))

	return nil
}
