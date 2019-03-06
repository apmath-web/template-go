package routing

import (
	"github.com/apmath-web/template-go/actions"
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/hello_world", actions.HelloWorldHandler)
	}
	return router
}
