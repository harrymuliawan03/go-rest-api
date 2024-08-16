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

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := authApi{authService: authService}

	app.Post("/auth/login", aa.Login)
}

func (aa authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseErrorData("Validation error", &fails))
	}

	res, err := aa.authService.Login(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess("Succesfully login", &res))
}
