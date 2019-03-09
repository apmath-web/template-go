package actions

import (
	"github.com/apmath-web/template-go/Application/viewModels"
	"github.com/apmath-web/template-go/Infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HelloWorldHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, viewModels.GenHelloWorldViewModel("wrong id", "400"))
	}
	hw := repositories.Repo.GetModel(id)
	helloWorldViewModel := viewModels.GenHelloWorldViewModel(hw.GetMessage(), "200")
	c.JSON(200, helloWorldViewModel)
}
