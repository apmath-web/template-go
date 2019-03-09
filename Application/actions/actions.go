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
		c.JSON(http.StatusBadRequest, viewModels.GenHelloWorldViewModel("Wrong id", "400"))
		return
	}
	hw := repositories.Repo.GetModel(id)
	if hw == nil {
		c.JSON(http.StatusNotFound, viewModels.GenHelloWorldViewModel("No model", "404"))
		return
	}
	helloWorldViewModel := viewModels.GenHelloWorldViewModel(hw.GetMessage(), "200")
	c.JSON(http.StatusOK, helloWorldViewModel)
}
