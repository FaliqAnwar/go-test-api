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

func (uc *customerUsecase) GetCustomerByID(ctx context.Context, in *model.RequestGetCustomerByID) (*model.Customer, error) {
	customer, err := uc.ucs.customerRepository.GetCustomerByID(ctx, in)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return customer, nil
}

func (uc *customerUsecase) NewCustomer(ctx context.Context, in *model.Customer) (*model.Customer, error) {
	//check if customer exist
	checkCustomer, err := uc.ucs.customerRepository.GetCustomerByID(ctx, &model.RequestGetCustomerByID{ID: in.ID})
	if checkCustomer != nil {
		return nil, fmt.Errorf("Customer %v is exist, please use PUT methode to update", checkCustomer.ID)
	}

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
		return nil, fmt.Errorf("%w", err)
	}

	return customers, nil
}

func (uc *customerUsecase) UpdateCustomer(ctx context.Context, in *model.RequestUpdateCustomer) (*model.Customer, error) {
	checkCustomer, err := uc.ucs.customerRepository.GetCustomerByID(ctx, &model.RequestGetCustomerByID{ID: in.ID})
	if checkCustomer == nil {
		return nil, fmt.Errorf("%w", err)
	}

	customer, err := uc.ucs.customerRepository.UpdateCustomer(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}

	if customer == checkCustomer {
		return nil, fmt.Errorf("record not change")
	}

	return customer, nil
}
