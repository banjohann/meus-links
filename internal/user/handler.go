package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Handlers are only responsible for:
// - Get the Request body, Request params and query params
// - Write the response data

type UserHandler struct {
	service *UserService
}

func NewHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func writeError(w http.ResponseWriter, msg string, err error) {
	http.Error(w, msg+err.Error(), http.StatusBadRequest)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserReq
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		writeError(w, "Failed to create user: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewCreateUserResp(*user))
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateUser(user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
		router.Get("/{userID}", h.getUser)
		router.Put("/", h.updateUser)
		router.Delete("/{userID}", h.deleteUser)
	}
}
