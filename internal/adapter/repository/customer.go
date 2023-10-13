package repository

import (
	"context"
	"errors"
	"fmt"
	"go-test-api/internal/adapter/repository/dto"
	"go-test-api/internal/model"
	"go-test-api/internal/port"

	"gorm.io/gorm"
)

type customer client

var _ port.CustomerRepository = (*customer)(nil)

func (c *Client) GetCustomerRepository() port.CustomerRepository {
	return c.customer
}

func (cr *customer) GetCustomerByID(ctx context.Context, in *model.RequestGetCustomerByID) (*model.Customer, error) {
	qdb := cr.c.db.WithContext(ctx).Model(&dto.CustomerDto{})
	qdb = qdb.Where(fmt.Sprintf("%s.%s = ?", dto.CustomerTableName, dto.CustomerIDField), in.ID)

	var customer dto.CustomerDto
	if err := qdb.First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return customer.ToEntity(), nil
}

func (cr *customer) Customer(ctx context.Context, in *model.Customer) (*model.Customer, error) {
	dto := dto.CustomerDto{}.FromModel(in)
	if err := cr.c.db.WithContext(ctx).Save(&dto).Error; err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}

func (cr *customer) GetCustomers(ctx context.Context) (model.Customers, error) {
	qdb := cr.c.db.WithContext(ctx).Model(&dto.CustomersDto{})

	var customers dto.CustomersDto
	if err := qdb.
		Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers.ToEntities(), nil
}

func (cr *customer) UpdateCustomer(ctx context.Context, in *model.RequestUpdateCustomer) (*model.Customer, error) {
	dto := dto.CustomerDto{}.FromModel(&model.Customer{
		ID:      in.ID,
		Name:    in.Name,
		Age:     in.Age,
		Address: in.Address,
		Salary:  in.Salary,
		Sector:  in.Sector,
	})

	if err := cr.c.db.WithContext(ctx).Save(&dto).Error; err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}
