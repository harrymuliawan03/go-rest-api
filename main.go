package main

import (
	"log"

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

	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository := repository.NewBookStock(dbConnection)
	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)

	bookService := service.NewBook(bookRepository, bookStockRepository)
	bookStockService := service.NewBookStock(bookStockRepository, bookRepository)
	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewBook(app, bookService, jwtMidd)
	api.NewBookStock(app, bookStockService, jwtMidd)
	api.NewAuth(app, authService)

	log.Fatal(app.Listen(":" + cnf.Server.Port))
}
