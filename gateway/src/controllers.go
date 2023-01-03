package src

import (
	"github.com/labstack/echo/v4"
)

type RouterController struct {
	UseCase SendRequestUseCase
}

func (ctrl *RouterController) Get(c echo.Context) error {
	response, err := ctrl.UseCase.Invoke(c.Request())
	if err != nil {
		return err
	}
	return c.JSON(response.StatusCode, response)
}

func (ctrl *RouterController) Post(c echo.Context) error {
	return nil
}
