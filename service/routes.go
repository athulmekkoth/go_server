package user

import (
	"net/http"

	"github.com/athulmekkoth/go_server.git/types"
	"github.com/athulmekkoth/go_server.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handelLogin).Methods("POST")
	router.HandleFunc("/register", h.handelRegister).Methods("POST")
}

func (h *Handler) handelLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.parseJson(r.body,payload)
	// var payload types.RegisterUserPayload

}
func (h *Handler) handelRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.parseJson(r,payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,err)
	}

}
