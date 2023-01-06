package src

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RouterController struct {
	UseCase GetRouteUseCase
}

func (ctrl *RouterController) Get(c echo.Context) error {
	_, err := ctrl.UseCase.Invoke(c.Path())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "")
}
