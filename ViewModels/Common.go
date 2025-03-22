package ViewModels

type CommonViewModel struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}