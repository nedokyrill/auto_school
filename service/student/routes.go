package student

import (
	"github.com/gorilla/mux"
	"newWebServer/types"
)

type Handler struct {
	store types.StudentStore
}

func NewHandler(store types.StudentStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

}
