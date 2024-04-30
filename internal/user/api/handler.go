package user_api

import (
	"encoding/json"
	"net/http"

	api "github.com/JohannBandelow/meus-links-go/internal/api"
	"github.com/JohannBandelow/meus-links-go/internal/user/service"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
	"github.com/go-chi/chi"
)

type UserHandler struct {
	service *service.UserService
}

func New(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req service.CreateUserReq
	encoder := json.NewEncoder(w)
	utils.DecodeBody(r, encoder, &req)

	user, err := h.service.CreateUser(req)
	if err != nil {
		encoder.Encode(api.ErrorBadRequest(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	//test
	encoder.Encode(user)
}

func (h *UserHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	var req service.LoginUserReq
	encoder := json.NewEncoder(w)

	utils.DecodeBody(r, encoder, &req)

	user, err := h.service.LoginUser(req)
	if err != nil {
		encoder.Encode(api.ErrorBadRequest(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(*user)
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	// var user user.User
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }

	// err = h.service.UpdateUser(user)
	// if err != nil {
	// 	http.Error(w, "Failed to update user", http.StatusBadRequest)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	// Retrieve the user from the service
	user, err := h.service.GetUser(userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	// Delete the user from the service
	err := h.service.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Failed to delete user with id: "+userID, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) LoadUserRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/", h.createUser)
		router.Post("/login", h.loginUser)
		router.Get("/{userID}", h.getUser)
		router.Put("/", h.updateUser)
		router.Delete("/{userID}", h.deleteUser)
	}
}
