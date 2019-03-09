package actions

import (
	"github.com/apmath-web/template-go/Application/viewModels"
	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	hw := viewModels.GenHelloWorld("hello world", "200 OK")
	c.JSON(200, hw)
}
