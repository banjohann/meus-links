package handler

import (
	"net/http"
)

type CreateUserReq struct{}

type CreateUserResp struct{}

func (router *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
}
