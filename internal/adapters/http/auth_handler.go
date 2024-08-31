package api

import (
	"encoding/json"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/service"
)

type AuthController struct {
	service service.AuthService
}

func (h *AuthController) login(w http.ResponseWriter, r *http.Request) {
	var req service.LoginCmd

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	user, err := h.service.Login(req)
	if err != nil {
		ErrorBadRequest("Erro ao realizar login", err.Error(), w, r)
		return
	}

	JSONResponse(w, http.StatusOK, "Login realizado com sucesso", user)
}
