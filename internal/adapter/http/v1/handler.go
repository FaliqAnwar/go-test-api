package v1

import (
	"go-test-api/internal/port"

	"github.com/labstack/echo/v4"
)

type customerHandler struct {
	customerUsecase port.CustomerUsecase
}

// New document handler will initialize the document/ resources endpoint
func New(app *echo.Group, uc port.Usecases) {
	c := customerHandler{
		customerUsecase: uc.GetCustomerUsecase(),
	}

	app.POST("/customer", c.customer())
	app.GET("/customer", c.getCustomers())
}
