package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrymuliawan03/go-rest-api/internal/config"
	"github.com/harrymuliawan03/go-rest-api/internal/connection"
	"github.com/harrymuliawan03/go-rest-api/internal/repository"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
