package helper

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Response struct {
	Status      string      `json:"status" example:"OK/ERROR"`
	Description string      `json:"description" example:"SUCCESS/ERROR"`
	Data        interface{} `json:"data,omitempty"`
	List        interface{} `json:"list,omitempty"`
}

func ResponseError(c echo.Context, status int, err error) error {
	res := Response{
		Status:      "ERROR",
		Description: err.Error(),
	}
	return c.JSON(status, res)
}

func ResponseSuccess(c echo.Context, r interface{}) error {
	res := Response{
		Status:      "OK",
		Description: "SUCCESS",
		Data:        r,
	}
	return c.JSON(http.StatusOK, res)
}

func ResponseSuccessList(c echo.Context, r interface{}) error {
	res := Response{
		Status:      "OK",
		Description: "SUCCESS",
		List:        r,
	}
	return c.JSON(http.StatusOK, res)
}
