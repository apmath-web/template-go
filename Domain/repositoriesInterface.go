package Domain

import "github.com/apmath-web/template-go/Infrastructure/applicationModels"

type repositoryInterface interface {
	GetModel(id int)
	PutModel(model applicationModels.HelloWorld) int
}
