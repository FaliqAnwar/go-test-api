package v1

import (
	"go-test-api/internal/adapter/http/helper"
	"go-test-api/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	DoPostNewCustomer struct {
		ID      int32  `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Age     int16  `json:"age,omitempty"`
		Address string `json:"address,omitempty"`
		Salary  int64  `json:"salary,omitempty"`
		Sector  int32  `json:"sector,omitempty"`
	}
)

func (h customerHandler) customer() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := DoPostNewCustomer{}

		if err := c.Bind(&req); err != nil {
			return helper.ResponseError(c, 401, err)
		}

		in := &model.Customer{
			ID:      req.ID,
			Name:    req.Name,
			Age:     req.Age,
			Address: req.Address,
			Salary:  req.Salary,
			Sector:  req.Sector,
		}

		out, err := h.customerUsecase.NewCustomer(ctx, in)
		if err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		res := DoPostNewCustomer(*out)

		return helper.ResponseSuccess(c, res)
	}
}

func (h customerHandler) getCustomers() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		out, err := h.customerUsecase.GetCustomers(ctx)
		if err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return helper.ResponseSuccessList(c, out)
	}
}
