package usecase

import (
	"context"
	"go-test-api/internal/model"
	"go-test-api/internal/port"
)

type customerUsecase usecase

var _ port.CustomerUsecase = (*customerUsecase)(nil)

func (ucs *Usecases) GetCustomerUsecase() port.CustomerUsecase {
	return ucs.customerUsecase
}

func (uc *customerUsecase) NewCustomer(ctx context.Context, in *model.Customer) (*model.Customer, error) {

	return &model.Customer{}, nil
}
