package usecase

import (
	"context"
	"fmt"
	"go-test-api/internal/model"
	"go-test-api/internal/port"
)

type customerUsecase usecase

var _ port.CustomerUsecase = (*customerUsecase)(nil)

func (ucs *Usecases) GetCustomerUsecase() port.CustomerUsecase {
	return ucs.customerUsecase
}

func (uc *customerUsecase) NewCustomer(ctx context.Context, in *model.Customer) (*model.Customer, error) {
	customer, err := uc.ucs.customerRepository.Customer(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	out := customer

	return out, nil
}

func (uc *customerUsecase) GetCustomers(ctx context.Context) (model.Customers, error) {
	customers, err := uc.ucs.customerRepository.GetCustomers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}

	return customers, nil
}
