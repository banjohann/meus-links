package api

import (
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/usecase/user"
	"github.com/go-chi/chi"
)

type UserController struct {
	CriaUsuarioUseCase     user.CriarUsuarioUseCase
	AtualizaUsuarioUseCase user.AtualizaUsuarioUseCase
	RemoveUsuarioUseCase   user.RemoveUsuarioUseCase
	GetUsuarioUseCase      user.GetUsuarioByIdUseCase
}

func (h *UserController) LoadRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/", h.create)
		router.Get("/{id}", h.getByID)
		router.Post("/{id}", h.update)
		router.Delete("/{id}", h.delete)
	}
}

func (h *UserController) create(w http.ResponseWriter, r *http.Request) {
	var req user.CriarUsuarioCmd
	err := DecodeBody(r, &req)
	if err != nil {
		ErrorBadRequest("Corpo da requisição em formato inválido", err.Error(), w, r)
		return
	}

	user, err := h.CriaUsuarioUseCase.Handle(req)
	if err != nil {
		ErrorBadRequest("Erro ao criar usuário", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusCreated, "Usuário criado com sucesso", user)
}

func (h *UserController) update(w http.ResponseWriter, r *http.Request) {
	var cmd user.AtualizaUsuarioCmd
	err := DecodeBody(r, &cmd)

	if err != nil {
		ErrorBadRequest("Corpo da requisição em formato inválido", err.Error(), w, r)
		return
	}

	userID := chi.URLParam(r, "id")
	cmd.ID = userID

	resp, err := h.AtualizaUsuarioUseCase.Handle(cmd)
	if err != nil {
		ErrorBadRequest("Erro ao atualizar usuário.", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Usuário atualizado com sucesso", resp)
}

func (h *UserController) getByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	user, err := h.GetUsuarioUseCase.Handle(userID)
	if err != nil {
		ErrorBadRequest("Erro ao buscar usuário por id.", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Usuário encontrado com sucesso", user)
}

func (h *UserController) delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	err := h.RemoveUsuarioUseCase.Handle(userID)
	if err != nil {
		ErrorBadRequest("Erro ao remover usuário.", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusNoContent, "Usuário removido com sucesso", userID)
}
