package handler

import (
	"net/http"

	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/infra/http/handler/web/presenter"
	"github.com/tomocy/archs/infra/http/view"
	"github.com/tomocy/archs/usecase"
)

func New(
	view view.View,
	userUsecase usecase.UserUsecase,
) *Handler {
	return &Handler{
		view:        view,
		userHandler: newUserHandler(view, userUsecase),
	}
}

type Handler struct {
	view        view.View
	userHandler *userHandler
}

func (h *Handler) ShowUserRegistrationForm(w http.ResponseWriter, r *http.Request) {
	h.userHandler.showRegistrationForm(w, r)
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	h.userHandler.registerUser(w, r)
}

func httpController(r *http.Request) *controller.HTTPController {
	return controller.NewHTTPController(r)
}

func webPresenter(view view.View, w http.ResponseWriter, r *http.Request) *presenter.Presenter {
	return presenter.New(view, w, r)
}
