package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
	"github.com/harrymuliawan03/go-rest-api/internal/util"
)

type bookApi struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, bs domain.BookService, jwtMidd fiber.Handler) {
	ba := bookApi{bookService: bs}

	app.Get("/books", jwtMidd, ba.Index)
	app.Get("/books/:id", jwtMidd, ba.Show)
	app.Post("/books", jwtMidd, ba.Create)
	app.Put("/books/:id", jwtMidd, ba.Update)
	app.Delete("/books/:id", jwtMidd, ba.Delete)
}

func (ba *bookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ba.bookService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess("Successfully get books", &res))
}

func (ba *bookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ba.bookService.Show(c, ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}
	if res.Id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError("Book not found"))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess("Succesfully get book", &res))
}

func (ba *bookApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.ResponseErrorData("Validation error", &fails))
	}

	err := ba.bookService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.ResponseSuccess[any]("Successfully create book", nil, 201))
}

func (ba *bookApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.ResponseErrorData("Validation error", &fails))
	}

	req.Id = ctx.Params("id")
	err := ba.bookService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess[any]("Successfully update book", nil))
}

func (ba *bookApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	err := ba.bookService.Delete(c, ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess[any]("Successfully delete customer", nil))
}
