package user_api

import (
	"encoding/json"
	"net/http"

	api "github.com/JohannBandelow/meus-links-go/internal/api"
	"github.com/JohannBandelow/meus-links-go/internal/user/service"
	"github.com/go-chi/chi"
)

type UserController struct {
	service *service.UserService
}

func New(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (h *UserController) create(w http.ResponseWriter, r *http.Request) {
	var req service.CriarUsuarioCmd
	encoder := json.NewEncoder(w)
	api.DecodeBody(r, &req)

	user, err := h.service.CriaUsuario(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao criar usuário", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(user)
}

func (h *UserController) login(w http.ResponseWriter, r *http.Request) {
	var req service.LoginCmd

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	user, err := h.service.Login(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*user)
}

func (h *UserController) update(w http.ResponseWriter, r *http.Request) {
	var cmd service.AtualizaUsuarioCmd
	err := json.NewDecoder(r.Body).Decode(&cmd)

	if err != nil {
		api.ErrorInternal("Erro ao atualizar usuário.", err.Error(), w, r)
		return
	}

	userID := chi.URLParam(r, "userID")
	cmd.ID = userID

	err = h.service.AtualizaUsuario(cmd)
	if err != nil {
		api.ErrorBadRequest("Erro ao atualizar usuário.", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(service.UpdateUserResponse{ID: cmd.ID})
}

func (h *UserController) getByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	user, err := h.service.GetUsuarioByID(userID)
	if err != nil {
		api.ErrorBadRequest("Erro ao buscar usuário por id.", err.Error(), w, r)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserController) delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	err := h.service.RemoveUsuario(userID)
	if err != nil {
		http.Error(w, "Failed to delete user with id: "+userID, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserController) LoadUserRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/", h.create)
		router.Post("/login", h.login)
		router.Get("/{userID}", h.getByID)
		router.Post("/{userID}", h.update)
		router.Delete("/{userID}", h.delete)
	}
}
