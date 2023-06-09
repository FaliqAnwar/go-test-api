package repository

import (
	"context"
	"go-test-api/internal/adapter/repository/dto"
	"go-test-api/internal/model"
	"go-test-api/internal/port"
)

type customer client

var _ port.CustomerRepository = (*customer)(nil)

func (c *Client) GetCustomerRepository() port.CustomerRepository {
	return c.customer
}

func (cr *customer) Customer(ctx context.Context, in *model.Customer) (*model.Customer, error) {
	dto := dto.CustomerDto{}.FromModel(in)
	if err := cr.c.db.WithContext(ctx).Save(&dto).Error; err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}

func (cr *customer) GetCustomers(ctx context.Context) (model.Customers, error) {
	qdb := cr.c.db.WithContext(ctx).Model(&dto.CustomerDto{})

	var customers dto.CustomersDto
	if err := qdb.
		Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers.ToEntities(), nil
}
