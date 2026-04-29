package user

import (
	"encoding/json"
	"net/http"

	"github.com/noyandey/go-ddd-starter/domain"
	"github.com/noyandey/go-ddd-starter/utils"
)

type CreateUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Role      string `json:"role"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreateUserRequest

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newUser)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	user, err := h.svc.Create(domain.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  hashedPassword,
		Role:      newUser.Role,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "User created successfully", user, http.StatusCreated)
}
