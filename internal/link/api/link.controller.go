package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/api"
	link_service "github.com/JohannBandelow/meus-links-go/internal/link/service"
	"github.com/go-chi/chi"
)

type LinkController struct {
	service *link_service.LinkService
}

func New(service *link_service.LinkService) *LinkController {
	return &LinkController{
		service: service,
	}
}

func (h *LinkController) LoadControllerRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/", h.create)
		router.Delete("/{id}", h.delete)
		router.Get("/", h.getAll)
		router.Put("/{id}", h.update)
		router.Get("/{id}", h.getById)
	}
}

func (h *LinkController) create(w http.ResponseWriter, r *http.Request) {
	var req link_service.CriaLinkCmd
	encoder := json.NewEncoder(w)
	err := api.DecodeBody(r, &req)
	if err != nil {
		api.ErrorInternal("Erro ao criar link", err.Error(), w, r)
		return
	}

	link, err := h.service.CriaLink(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao criar link", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(link)
}

func (h *LinkController) delete(w http.ResponseWriter, r *http.Request) {
	linkID := chi.URLParam(r, "id")
	err := h.service.RemoverLink(linkID)
	if err != nil {
		api.ErrorBadRequest("Erro ao deletar link", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LinkController) getAll(w http.ResponseWriter, r *http.Request) {
	var req link_service.GetAllRequestFilter
	encoder := json.NewEncoder(w)
	links, err := h.service.GetAll(req)
	if err != nil {
		api.ErrorBadRequest("Erro ao deletar link", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(links)
}

func (h *LinkController) update(w http.ResponseWriter, r *http.Request) {
	var req link_service.AtualizaLinkCmd
	linkId := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	err := api.DecodeBody(r, &req)
	if err != nil {
		api.ErrorBadRequest("Erro ao atualizar link", err.Error(), w, r)
		return
	}

	link, err := h.service.AtualizaLink(linkId, req)
	if err != nil {
		api.ErrorBadRequest("Erro ao atualizar link", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(link)
}

func (h *LinkController) getById(w http.ResponseWriter, r *http.Request) {
	linkID := chi.URLParam(r, "id")
	encoder := json.NewEncoder(w)
	link, err := h.service.GetById(linkID)
	if err != nil {
		errMsg := fmt.Sprintf("Link não encontrado com o id: %s", linkID)
		api.ErrorBadRequest(errMsg, err.Error(), w, r)
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(link)
}
