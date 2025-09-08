package auth

import (
	"os"

	"github.com/absorkun/darinol/response"
	jwtware "github.com/absorkun/darinol/temporary/jwt"
	"github.com/gofiber/fiber/v3"
)

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
	e.Router.Get("/session", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("KEY"))},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return response.BadRequest(c, err.Error())
		},
	}), e.Handler.Session)
}
