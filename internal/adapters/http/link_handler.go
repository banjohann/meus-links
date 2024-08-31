package api

import (
	"fmt"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/usecase/link"
	"github.com/go-chi/chi"
)

type LinkController struct {
	CriaLinkUseCase     link.CriaLinkUseCase
	AtualizaLinkUseCase link.AtualizaLinkUseCase
	RemoverLinkUseCase  link.RemoverLinkUseCase
	GetByIdUseCase      link.GetByIdLinkUseCase
	GetAllUseCase       link.GetAllLinkUseCase
}

func (h *LinkController) LoadRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/", h.getAll)
		router.Post("/", h.create)
		router.Put("/{id}", h.update)
		router.Get("/{id}", h.getById)
		router.Delete("/{id}", h.delete)
	}
}

func (h *LinkController) create(w http.ResponseWriter, r *http.Request) {
	var req link.CriaLinkCmd
	err := DecodeBody(r, &req)
	if err != nil {
		ErrorInternal("Erro ao criar link", err.Error(), w, r)
		return
	}

	link, err := h.CriaLinkUseCase.Handle(req)
	if err != nil {
		ErrorBadRequest("Erro ao criar link", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusCreated, "Link criado com sucesso", link)
}

func (h *LinkController) delete(w http.ResponseWriter, r *http.Request) {
	linkID := chi.URLParam(r, "id")
	err := h.RemoverLinkUseCase.Handle(linkID)
	if err != nil {
		ErrorBadRequest("Erro ao deletar link", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Link deletado com sucesso", nil)
}

func (h *LinkController) update(w http.ResponseWriter, r *http.Request) {
	var req link.AtualizaLinkCmd
	linkId := chi.URLParam(r, "id")
	err := DecodeBody(r, &req)
	if err != nil {
		ErrorBadRequest("Erro ao atualizar link", err.Error(), w, r)
		return
	}

	link, err := h.AtualizaLinkUseCase.Handle(linkId, req)
	if err != nil {
		ErrorBadRequest("Erro ao atualizar link", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Link atualizado com sucesso", link)
}

func (h *LinkController) getById(w http.ResponseWriter, r *http.Request) {
	linkID := chi.URLParam(r, "id")
	link, err := h.GetByIdUseCase.Handle(linkID)
	if err != nil {
		errMsg := fmt.Sprintf("Link não encontrado com o id: %s", linkID)
		ErrorBadRequest(errMsg, err.Error(), w, r)
	}

	JSONResponse(w, http.StatusOK, "Link encontrado com sucesso", link)
}

func (h *LinkController) getAll(w http.ResponseWriter, r *http.Request) {
	var req link.GetAllRequestFilter
	links, err := h.GetAllUseCase.Handle(req)
	if err != nil {
		ErrorBadRequest("Erro ao deletar link", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Links encontrados com sucesso", links)
}
