package user_api

import (
	"encoding/json"
	"net/http"

	api "github.com/JohannBandelow/meus-links-go/internal/api"
	"github.com/JohannBandelow/meus-links-go/internal/user/service"
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
	var req service.CreateUserCmd
	encoder := json.NewEncoder(w)
	api.DecodeBody(r, &req)

	user, err := h.service.CreateUser(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao criar usuário", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(user)
}

func (h *UserHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	var req service.LoginUserCmd

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	user, err := h.service.LoginUser(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*user)
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var cmd service.UpdateUserCmd
	err := json.NewDecoder(r.Body).Decode(&cmd)

	if err != nil {
		api.ErrorInternal("Erro ao atualizar usuário.", err.Error(), w, r)
		return
	}

	userID := chi.URLParam(r, "userID")
	cmd.ID = userID

	err = h.service.UpdateUser(cmd)
	if err != nil {
		api.ErrorBadRequest("Erro ao atualizar usuário.", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(service.UpdateUserResponse{ID: cmd.ID})
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		api.ErrorBadRequest("Erro ao buscar usuário por id.", err.Error(), w, r)
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
		router.Post("/{userID}", h.updateUser)
		router.Delete("/{userID}", h.deleteUser)
	}
}
