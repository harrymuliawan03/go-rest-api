package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/harrymuliawan03/go-rest-api/dto"
	"github.com/harrymuliawan03/go-rest-api/internal/api"
	"github.com/harrymuliawan03/go-rest-api/internal/config"
	"github.com/harrymuliawan03/go-rest-api/internal/connection"
	"github.com/harrymuliawan03/go-rest-api/internal/repository"
	"github.com/harrymuliawan03/go-rest-api/internal/service"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMidd := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ResponseError("Unauthorized"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)

	bookRepository := repository.NewBook(dbConnection)
	bookService := service.NewBook(bookRepository)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewAuth(cnf, userRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewBook(app, bookService, jwtMidd)
	api.NewAuth(app, authService)

	_ = app.Listen(":" + cnf.Server.Port)
}
