package main

import (
	"log"
	"os"

	"github.com/absorkun/darinol/auth"
	"github.com/absorkun/darinol/database"
	"github.com/absorkun/darinol/temporary/swagger"
	"github.com/absorkun/darinol/todo"
	"github.com/absorkun/darinol/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	v.validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		var role = fl.Field().String()
		return role == "user" || role == "admin"
	})
	return v.validate.Struct(out)
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Format input: Bearer (access_token)
func main() {
	var db = database.New()
	var app = fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
	}))
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		Title:    "Swagger OpanAPI Documentation | darinol",
		FilePath: "./docs/swagger.json",
	}))
	app.Get("", func(c fiber.Ctx) error {
		return c.Redirect().To("/docs")
	})
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	var usersV1 = app.Group("/api/v1/users")
	var userHandler = user.NewHandler(db)
	var userEndpoint = user.NewEndpoint(usersV1, *userHandler)
	userEndpoint.Run()

	var todoV1 = app.Group("/api/v1/todos")
	var todoHandler = todo.NewHandler(db)
	var todoEndpoint = todo.NewEndpoint(todoV1, *todoHandler)
	todoEndpoint.Run()

	var authV1 = app.Group("/api/v1/auth")
	var authHandler = auth.NewHandler(db)
	var authEndpoint = auth.NewEndpoint(authV1, *authHandler)
	authEndpoint.Run()

	var port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	var host = os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	if err := app.Listen(host + ":" + port); err != nil {
		log.Fatal(err.Error())
	}

}
