package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type RouterController struct {
	GetServerUseCase
}

func (ctrl RouterController) Proxy(c *gin.Context) {
	path := c.Param("slug")

	// Get server specs
	server, err := ctrl.GetServerUseCase.Invoke(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
		return
	}

	// Get a new url
	// TODO: Create a use case to get a path without the first string
	proxyUrl, err := url.Parse(fmt.Sprintf("%s%s", server.Host, "random.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem to parse url"})
		return
	}

	// Make a new request
	// TODO: Remove this dependency hard coded
	reverseProxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	reverseProxy.Director = func(request *http.Request) {
		request.Header = c.Request.Header
		request.Host = proxyUrl.Host
		request.URL = proxyUrl
	}
	reverseProxy.ServeHTTP(c.Writer, c.Request)
}
