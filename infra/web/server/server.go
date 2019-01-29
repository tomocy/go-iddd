package server

import (
	"net/http"

	"github.com/tomocy/archs/infra/web/handler"
	"github.com/tomocy/chi"
)

type Server interface {
	http.Handler
	RegisterRoute(h handler.Handler)
}

func NewServer() Server {
	return newChiServer()
}

type chiServer struct {
	router chi.Router
}

func newChiServer() *chiServer {
	return &chiServer{
		router: chi.NewRouter(),
	}
}

func (s chiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s chiServer) RegisterRoute(h handler.Handler) {
	s.router.Route("/users", func(r chi.Router) {
		r.Post("/", h.RegisterUser)
	})
}