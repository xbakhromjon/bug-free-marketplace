package controller

import (
	"encoding/json"
	jwt "golang-project-template/internal/pkg/jwt"
	"golang-project-template/internal/users/app"
	"golang-project-template/internal/users/domain"
	"net/http"
)

type loginRequest struct {
	phoneNumber string
	password    string
}

type UserController struct {
	userUsecase app.UserUsecase
}

func NewUserController(userUsecase app.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (c *UserController) RegisterAdminUserHandler(w http.ResponseWriter, r *http.Request) {
	var newAdmin domain.NewUser

	//decoding requested body to go object
	err := json.NewDecoder(r.Body).Decode(&newAdmin)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	//creating admin user in repo layer
	id, err := c.userUsecase.RegisterAdmin(&newAdmin)
	if err != nil {
		http.Error(w, "failed to create an admin user: "+err.Error(), http.StatusInternalServerError)
	}

	//writing to the HEADER
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (c *UserController) RegisterCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var newCustomer domain.NewUser

	//decoding requested body to go object
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	//creating customer user in repo layer
	id, err := c.userUsecase.RegisterCustomer(&newCustomer)
	if err != nil {
		http.Error(w, "failed to create customer user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//writing to the HEADER
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (c *UserController) RegisterMerchantHandler(w http.ResponseWriter, r *http.Request) {
	var newMerchant domain.NewUser

	//decoding requested body to go object
	err := json.NewDecoder(r.Body).Decode(&newMerchant)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	//creating customer user in repo layer
	id, err := c.userUsecase.RegisterMerchantUser(&newMerchant)
	if err != nil {
		http.Error(w, "failed to create merchant user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//writing to the HEADER
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (c *UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req loginRequest

	//decoding requested body to go object
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	passed, err := c.userUsecase.LoginUser(req.phoneNumber, req.password)
	if err != nil {
		http.Error(w, "Invalid password or phone number", http.StatusInternalServerError)
		return
	}

	if passed {
		token, err := jwt.CreateToken(req.phoneNumber)
		if err != nil {
			http.Error(w, "phone number not found: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"access_token": token,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}

}
