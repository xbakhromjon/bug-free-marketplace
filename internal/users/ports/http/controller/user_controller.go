package controller

import (
	"encoding/json"
	"golang-project-template/internal/pkg/jwt"
	"golang-project-template/internal/users/app"
	"golang-project-template/internal/users/domain"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type userObject struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phone_number"`
	Role        string     `json:"role"`
	CreateAt    time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type loginRequest struct {
	PhoneNumber string
	Password    string
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
	var req loginRequest

	//decoding requested body to go object
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	passed, err := c.userUsecase.LoginUser(req.PhoneNumber, req.Password)
	if err != nil {
		http.Error(w, "Invalid password or phone number", http.StatusInternalServerError)
		return
	}

	if passed {
		token, err := jwt.CreateToken(req.PhoneNumber)
		if err != nil {
			http.Error(w, "phone number not found: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"access_token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}

}

func (c *UserController) GetUserByPhoneNumberHandler(w http.ResponseWriter, r *http.Request) {

	phoneNumber := chi.URLParam(r, "phone_number")
	var newUser userObject

	user, err := c.userUsecase.GetUserByPhoneNumber(phoneNumber)

	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newUser.Id = user.GetID()
	newUser.Name = user.GetName()
	newUser.PhoneNumber = user.GetPhoneNumber()
	newUser.Role = user.GetRole()
	newUser.CreateAt = user.GetCreatedAt()
	newUser.UpdatedAt = user.GetUpdatedAt()
	newUser.DeletedAt = user.GetDeletedAt()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newUser)

}
