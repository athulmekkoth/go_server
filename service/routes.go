package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handelLogin).Methods("POST")
	router.HandleFunc("/register", h.handelRegister).Methods("POST")
}

func (h *Handler) handelLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handelRegister(w http.ResponseWriter, r *http.Request) {

}
