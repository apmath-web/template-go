package repositories

import (
	"github.com/apmath-web/template-go/Domain"
)

type Repository struct {
	models         map[int]Domain.HelloWorldApplicatiomModel
	numberOfModels int
}

func GenRepository() *Repository {
	repo := &Repository{make(map[int]Domain.HelloWorldApplicatiomModel), 0}
	return repo
}

func (r *Repository) GetModel(id int) Domain.HelloWorldApplicatiomModel {
	helloWorld, ok := r.models[id]
	if ok {
		return helloWorld
	}
	return nil
}

func (r *Repository) PutModel(model Domain.HelloWorldApplicatiomModel) int {
	r.numberOfModels++
	r.models[r.numberOfModels] = model
	return r.numberOfModels
}
