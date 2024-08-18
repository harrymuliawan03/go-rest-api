package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
	"github.com/harrymuliawan03/go-rest-api/internal/util"
)

type BookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, bss domain.BookStockService, jwtMidd fiber.Handler) {
	bsa := BookStockApi{bookStockService: bss}

	app.Post("/book-stocks", jwtMidd, bsa.Create)
	app.Delete("/book-stocks", jwtMidd, bsa.Delete)
}

func (bsa BookStockApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	var statusCode int
	if err := ctx.BodyParser(&req); err != nil {
		statusCode = http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.ResponseError(err.Error(), statusCode))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode = http.StatusUnprocessableEntity
		return ctx.Status(statusCode).JSON(dto.ResponseErrorData("Validation failed", &fails, statusCode))
	}

	err := bsa.bookStockService.Create(c, req)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.ResponseError(err.Error(), statusCode))
	}

	statusCode = http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.ResponseSuccess[any]("Successfully create book stock", nil, statusCode))
}

func (bsa BookStockApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var statusCode int
	// ?=code=
	codeStr := ctx.Query("code")
	codes := strings.Split(codeStr, ";")
	if len(codes) < 1 || codeStr == "" {
		statusCode = http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.ResponseError("Parameter code wajib diisi", statusCode))
	}

	err := bsa.bookStockService.Delete(c, dto.DeleteBookStockRequest{Codes: codes})
	if err != nil {
		statusCode = http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.ResponseError(err.Error(), statusCode))
	}

	statusCode = http.StatusOK
	return ctx.Status(statusCode).JSON(dto.ResponseSuccess[any]("Delete book stock successfully", nil, statusCode))
}
