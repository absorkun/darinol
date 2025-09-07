package user

import "github.com/gofiber/fiber/v3"

type endpoint struct {
	Router  fiber.Router
	Handler handler
}

func NewEndpoint(router fiber.Router, handler handler) *endpoint {
	return &endpoint{router, handler}
}

func (e *endpoint) Run() {
	e.Router.Get("", e.Handler.GetAll)
	e.Router.Get("/:id", e.Handler.GetById)
	e.Router.Post("", e.Handler.Create)
	e.Router.Put("/:id", e.Handler.Update)
	e.Router.Delete("/:id", e.Handler.Delete)
}
