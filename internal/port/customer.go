package port

import (
	"context"
	"go-test-api/internal/model"
)

type CustomerUsecase interface {
	NewCustomer(ctx context.Context, in *model.Customer) (*model.Customer, error)
	GetCustomers(ctx context.Context) (model.Customers, error)
	GetCustomerByID(ctx context.Context, in *model.RequestGetCustomerByID) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, in *model.RequestUpdateCustomer) (*model.Customer, error)
}

type CustomerRepository interface {
	Customer(ctx context.Context, in *model.Customer) (*model.Customer, error)
	GetCustomers(ctx context.Context) (model.Customers, error)
	GetCustomerByID(ctx context.Context, in *model.RequestGetCustomerByID) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, in *model.RequestUpdateCustomer) (*model.Customer, error)
}
