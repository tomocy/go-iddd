package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/web/validator"
)

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	controller controller.UserController
}

func NewUserHandler(controller controller.UserController) UserHandler {
	return &userHandler{
		controller: controller,
	}
}

func (h userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	validated, err := validator.ValidateToRegisterUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.controller.RegisterUser(validated.Email, validated.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "register user: {id: %s, email: %s}\n", resp.ID, resp.Email)
}