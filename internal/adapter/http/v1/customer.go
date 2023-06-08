package v1

import (
	"go-test-api/internal/adapter/http/helper"

	"github.com/labstack/echo/v4"
)

type (
	DoPostNewCustomer struct {
		ID      int32  `json:"id"`
		Name    string `json:"name"`
		Age     int16  `json:"age"`
		Address string `json:"address"`
		Salary  int64  `json:"salary"`
		Sector  int32  `json:"sector"`
	}
)

func (h customerHandler) customer() echo.HandlerFunc {
	return func(c echo.Context) error {
		//ctx := c.Request().Context()
		//
		//var n int
		//id, _ := fmt.Sscan(c.FormValue("id"), &n)
		//age, _ := fmt.Sscan(c.FormValue("age"), &n)
		//salary, _ := fmt.Sscan(c.FormValue("salary"), &n)
		//sector, _ := fmt.Sscan(c.FormValue("sector"), &n)
		//
		//in := model.Customer{
		//	ID:      int32(id),
		//	Name:    c.FormValue("name"),
		//	Age:     int16(age),
		//	Address: c.FormValue("address"),
		//	Salary:  int64(salary),
		//	Sector:  int32(sector),
		//}
		//
		//out, err := h.customerUsecase.NewCustomer(ctx, &in)
		//if err != nil {
		//	return helper.ResponseError(c, http.StatusInternalServerError, err)
		//}

		//res := DoPostNewCustomer(*out)
		res := "test hit api success"

		return helper.ResponseSuccess(c, res)
	}
}
