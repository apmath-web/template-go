package Domain

type repositoryInterface interface {
	GetModel(id int) HelloWorldApplicatiomModel
	PutModel(model HelloWorldApplicatiomModel) int
}
