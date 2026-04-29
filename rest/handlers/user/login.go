package user

import (
	"encoding/json"
	"net/http"

	"github.com/noyandey/go-ddd-starter/domain"
	"github.com/noyandey/go-ddd-starter/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string       `json:"accessToken"`
	RefreshToken *string      `json:"refreshToken"`
	User         *domain.User `json:"user"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&reqLogin)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	user, err := h.svc.Find(reqLogin.Email, reqLogin.Password)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if user == nil {
		utils.SendError(w, false, "Invalid email or password", nil, http.StatusBadRequest)
		return
	}

	accessToken, err := utils.CreateJwt(h.cfg.JwtSecretKey, utils.Payload{
		Sub:       int64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(
		w,
		true,
		"Login successful",
		LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: nil,
			User:         user,
		},
		http.StatusOK,
	)

}
