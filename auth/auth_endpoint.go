package auth

import "github.com/gofiber/fiber/v3"

type endpoint struct {
	Router  fiber.Router
	Handler handler
}

func NewEndpoint(router fiber.Router, handler handler) *endpoint {
	return &endpoint{router, handler}
}

func (e *endpoint) Run() {
	e.Router.Post("/login", e.Handler.Login)
	e.Router.Post("/register", e.Handler.Register)
}
