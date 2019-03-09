package applicationModels

type HelloWorld struct {
	Message string `json:"message"`
}

func (hw *HelloWorld) GetMessage() string {
	return hw.Message
}

func (hw *HelloWorld) SetMessage(message string) string {
	hw.Message = message
}
