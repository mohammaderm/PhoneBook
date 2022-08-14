package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mohammaderm/PhoneBook/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretkey = "mohammad123456789@"
	issue     = "127.0.1.1"
)

type JwtClaims struct {
	Email string `json:"email"`
	Id    uint   `json:"id"`
	jwt.StandardClaims
}

func (h *HandlerFields) Login(w http.ResponseWriter, r *http.Request) {

	var user models.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "can not parse values.", http.StatusBadRequest)
		return
	}
	founduser, err := h.repository.GetbyEmail_User(user.Email)
	log.Println(err)
	if err != nil {
		http.Error(w, "can not found user registered with this email.", http.StatusForbidden)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(founduser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Wrong password!", http.StatusForbidden)
		return
	}
	pairtoken, err := generatepairtoken(founduser.Email, founduser.Id)
	if err != nil {
		http.Error(w, "internal server.", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	jsonresp, _ := json.Marshal(pairtoken)
	w.Write(jsonresp)

}

func (h *HandlerFields) Register(w http.ResponseWriter, r *http.Request) {

	var userreq models.User
	err := json.NewDecoder(r.Body).Decode(&userreq)
	if err != nil {
		http.Error(w, "can not parse values.", http.StatusBadRequest)
		return
	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(userreq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "internal server.", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Username: userreq.Username,
		Email:    userreq.Email,
		Password: string(hashpass),
	}

	if err := h.repository.Create_User(user); err != nil {
		http.Error(w, "internal server,", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Sucsefully created User."
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)

}

func generatepairtoken(email string, id uint) (map[string]string, error) {
	// Access_token
	a_claims := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    issue,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, a_claims)
	a_token, err := jwtToken.SignedString([]byte(secretkey))
	if err != nil {
		return nil, err
	}
	// Refresh Token
	r_claims := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    issue,
		},
	}
	r_jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, r_claims)
	r_token, err := r_jwtToken.SignedString([]byte(secretkey))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  a_token,
		"refresh_token": r_token,
	}, nil

}
