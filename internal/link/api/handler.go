package api

import (
	"encoding/json"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/link/service"
)

type LinkHandler struct {
	service *service.LinkService
}

func New(service *service.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

func (h *LinkHandler) createLink(w http.ResponseWriter, r *http.Request) {
	var req service.CreateLinkCmd
	encoder := json.NewEncoder(w)
	DecodeBody(r, &req)

	link, err := h.service.CreateLink(req)
	if err != nil {
		ErrorBadRequest("Erro ao criar link", err.Error(), w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(link)
}
