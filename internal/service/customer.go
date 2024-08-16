package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(cr domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: cr,
	}
}

// Index implements domain.CustomerService.
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}
	customerData := make([]dto.CustomerData, 0, len(customers))
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			Id:   v.Id,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	generateId := time.Now().Unix()
	generateIdStr := strconv.FormatInt(generateId, 10)

	customer := domain.Customer{
		Id:        generateIdStr,
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return c.customerRepository.Save(ctx, &customer)
}

func (c *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.customerRepository.FindById(ctx, req.Id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("customer not found")
	}
	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return c.customerRepository.Update(ctx, &persisted)
}

func (c *customerService) Delete(ctx context.Context, id string) error {
	persisted, err := c.customerRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("Customer not found")
	}

	return c.customerRepository.Delete(ctx, id)
}

func (c *customerService) Show(ctx context.Context, id string) (cd *dto.CustomerData, err error) {
	persisted, err := c.customerRepository.FindById(ctx, id)

	if persisted.Id == "" {
		err = errors.New("Customer not found")
	} else {
		cd = &dto.CustomerData{
			Id:   persisted.Id,
			Code: persisted.Code,
			Name: persisted.Name,
		}
	}

	return
}
