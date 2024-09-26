package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"newWebServer/types"
	"newWebServer/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegicterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.loginHandler).Methods("POST")
	router.HandleFunc("/register", h.registerHandler).Methods("POST")
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	//1.принять json
	//2.проверить есть ли уже такой юзер
	//3.добавить юзера в бд
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	checkUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}

	if checkUser != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with %s email already exists", payload.Email))
	}

	h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.Nickname,
		Email:     payload.Email,
		Password:  payload.Password,
	})
}
