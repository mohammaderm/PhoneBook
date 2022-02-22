package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammaderm/PhoneBook/models"
	"github.com/mohammaderm/PhoneBook/repository"
)

type (
	HandlerFields struct {
		repository repository.MainRepository
	}
	HandlerInterface interface {
		// All Contact Handler
		Create_contact(w http.ResponseWriter, r *http.Request)
		Delete_contact(w http.ResponseWriter, r *http.Request)
		GetByName_contact(w http.ResponseWriter, r *http.Request)
		GetAll_contact(w http.ResponseWriter, r *http.Request)
		GetByCategory_contact(w http.ResponseWriter, r *http.Request)
		// --------------------------------------------------------------

		// All Category Handler
		Create_category(w http.ResponseWriter, r *http.Request)
		Delete_category(w http.ResponseWriter, r *http.Request)
		GetAll_category(w http.ResponseWriter, r *http.Request)
		// --------------------------------------------------------------

		// All User Authentication Handler
		Register(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
		// --------------------------------------------------------------

	}
)

func NewHandler(repo repository.MainRepository) HandlerInterface {
	return &HandlerFields{
		repository: repo,
	}
}

// --------------------------
// All Contact Handler
// --------------------------
func (h *HandlerFields) Create_contact(w http.ResponseWriter, r *http.Request) {

	userid := r.Context().Value("UserId")
	var contact models.ReqCreateContact
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&contact); err != nil {
		http.Error(w, "can not parse values.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	newcontact := &models.Cantact{
		Name:        contact.Name,
		LastName:    contact.LastName,
		Address:     contact.Address,
		PhoneNumber: contact.PhoneNumber,
		CategoryID:  contact.CategoryID,
		UserID:      userid.(uint),
	}

	if err := h.repository.Create_contact(newcontact); err != nil {
		http.Error(w, "internal server,", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Sucsefully created Contact."
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
}

func (h *HandlerFields) Delete_contact(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	if id == "" {
		http.Error(w, "You are missing ID parameter.", http.StatusBadRequest)
		return
	}
	userid := r.Context().Value("UserId")
	res, _ := h.repository.Delete_contact(id, strconv.FormatUint(uint64(userid.(uint)), 10))
	if !res {
		http.Error(w, "No contact found with this id.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "The contact has been deleted successfully!"
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
}

func (h *HandlerFields) GetByName_contact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		http.Error(w, "You are missing name parameter.", http.StatusBadRequest)
		return
	}
	userid := r.Context().Value("UserId")
	result, err := h.repository.GetByName_contact(name, strconv.FormatUint(uint64(userid.(uint)), 10))

	if err != nil {
		http.Error(w, "result not found.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonresp, _ := json.Marshal(result)
	w.Write(jsonresp)

}

func (h *HandlerFields) GetAll_contact(w http.ResponseWriter, r *http.Request) {
	userid := r.Context().Value("UserId")
	result, err := h.repository.GetAll_contact(strconv.FormatUint(uint64(userid.(uint)), 10))
	if err != nil {
		http.Error(w, "result not found.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonresp, _ := json.Marshal(result)
	w.Write(jsonresp)
}

func (h *HandlerFields) GetByCategory_contact(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	if id == "" {
		http.Error(w, "You are missing ID parameter.", http.StatusBadRequest)
		return
	}
	userid := r.Context().Value("UserId")
	res, err := h.repository.GetByCategoryId_contact(id, strconv.FormatUint(uint64(userid.(uint)), 10))
	fmt.Println(err)
	if err != nil {
		http.Error(w, "result not found.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonresp, _ := json.Marshal(res)
	w.Write(jsonresp)
}

// --------------------------------------------------------------

// --------------------------
// All Category Handler
// --------------------------
func (h *HandlerFields) Create_category(w http.ResponseWriter, r *http.Request) {

	userid := r.Context().Value("UserId")
	var category models.ReqCreateCategory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		http.Error(w, "can not parse values.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	newcategory := &models.Category{
		Name:   category.Name,
		UserID: userid.(uint),
	}

	if err := h.repository.Create_category(newcategory); err != nil {
		fmt.Println(err)
		http.Error(w, "internal server,", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Sucsefully created Category."
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
}

func (h *HandlerFields) Delete_category(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	if id == "" {
		http.Error(w, "invalid request params.", http.StatusBadRequest)
		return
	}
	userid := r.Context().Value("UserId")
	res, _ := h.repository.Delete_category(id, strconv.FormatUint(uint64(userid.(uint)), 10))
	if !res {
		http.Error(w, "No category found with this id.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "The category has been deleted successfully!"
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
}

func (h *HandlerFields) GetAll_category(w http.ResponseWriter, r *http.Request) {
	userid := r.Context().Value("UserId")
	result, err := h.repository.GetAll_category(strconv.FormatUint(uint64(userid.(uint)), 10))
	fmt.Println(err)
	if err != nil {
		http.Error(w, "result not found.", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonresp, _ := json.Marshal(result)
	w.Write(jsonresp)
}

// --------------------------------------------------------------
