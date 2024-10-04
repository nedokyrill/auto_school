package exam

import (
	"github.com/gorilla/mux"
	"newWebServer/types"
)

type Handler struct {
	store types.ExamStore
}

func NewHandler(store types.ExamStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

}
