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

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService, jwtMidd fiber.Handler) {
	ca := customerApi{customerService: customerService}

	app.Get("/customers", jwtMidd, ca.Index)
	app.Get("/customers/:id", jwtMidd, ca.Show)
	app.Post("/customers", jwtMidd, ca.Create)
	app.Put("/customers/:id", jwtMidd, ca.Update)
	app.Delete("/customers/:id", jwtMidd, ca.Delete)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.JSON(dto.ResponseSuccess("Succesfully get customers", &res))
}

func (ca customerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	cd, err := ca.customerService.Show(c, ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	if cd.Id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError("Customer not found"))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess("Successfully get customer", cd))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError(err.Error()))
	}
	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.ResponseErrorData("Validation error", &fails))
	}

	err := ca.customerService.Create(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.ResponseSuccess[any]("Succesfully create customer", nil, 201))
}

func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseError(err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ResponseErrorData("Validation error", &fails))
	}

	req.Id = ctx.Params("id")
	err := ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess[any]("Successfully update customer", nil))
}

func (ca customerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	err := ca.customerService.Delete(c, ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.ResponseSuccess[any]("Successfully delete customer", nil))
}
