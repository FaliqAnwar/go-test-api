package port

import (
	"context"
	"go-test-api/internal/model"
)

type CustomerUsecase interface {
	NewCustomer(ctx context.Context, in *model.Customer) (*model.Customer, error)
}

type CustomerRepository interface {
}
