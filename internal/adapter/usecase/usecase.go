package usecase

import (
	"go-test-api/internal/model"
	"go-test-api/internal/port"
)

type Usecases struct {
	customerUsecase    *customerUsecase
	customerRepository port.CustomerRepository

	common usecase
	config model.Config
}

type usecase struct {
	ucs *Usecases
}

func NewUsecases(config model.Config, customerRepository port.CustomerRepository) *Usecases {
	uc := &Usecases{
		config:             config,
		customerRepository: customerRepository,
	}

	uc.common.ucs = uc
	uc.customerUsecase = (*customerUsecase)(&uc.common)

	return uc
}
