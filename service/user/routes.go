package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"newWebServer/service/auth"
	"newWebServer/types/User"
	"newWebServer/utils"
	"os"
)

type Handler struct {
	store User.UserStore
}

func NewHandler(store User.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.loginHandler).Methods("POST")
	router.HandleFunc("/register", h.registerHandler).Methods("POST")
	//router.HandleFunc("/update", h.updateUserHandler).Methods("PATCH")
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var payload User.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate
	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//check user in db
	checkUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email"))
		return
	}

	if !auth.ComparePasswords(checkUser.UserDetails.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
		return
	}

	//jwt
	err = godotenv.Load()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	secret := []byte(os.Getenv("JWT_EXPIRATION"))
	token, err := auth.CreateJWT(secret, checkUser.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}

//func (h *Handler) updateUserHandler(w http.ResponseWriter, r *http.Request) {
//	var payload types.User
//	if err := utils.ParseJson(r, &payload); err != nil {
//		utils.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	_, err := h.store.GetUserByEmail(payload.UserDetails.Email)
//	if err != nil {
//		utils.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	hashPassword, err := auth.HashPassword(payload.UserDetails.Password)
//
//	if err != nil {
//		utils.WriteError(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	err = h.store.UpdateUser(types.User{
//		FirstName: payload.UserDetails.FirstName,
//		LastName:  payload.UserDetails.LastName,
//		Email:     payload.UserDetails.Email,
//		Password:  hashPassword,
//	})
//
//	if err != nil {
//		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error with updating user: ", err))
//		return
//	}
//
//	utils.WriteJson(w, http.StatusOK, payload)
//}

func (h *Handler) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	//1.принять json
	//2.проверить есть ли уже такой юзер
	//3.добавить юзера в бд
	var payload User.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate
	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	checkUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		hashPassword, err := auth.HashPassword(payload.Password)

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		err = h.store.CreateUser(User.User{
			UserDetails: User.UserDetails{
				FirstName: payload.FirstName,
				LastName:  payload.LastName,
				Email:     payload.Email,
				Password:  hashPassword,
			},
		})

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error with creating user: ", err))
			return
		}

		utils.WriteJson(w, http.StatusCreated, checkUser)
		return
	}

	utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s already exists", payload.Email))
}
