package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterController struct {
	UseCase GetRouteUseCase
}

func (ctrl *RouterController) Proxy(c *gin.Context) {
	_, err := ctrl.UseCase.Invoke(c.FullPath())
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
