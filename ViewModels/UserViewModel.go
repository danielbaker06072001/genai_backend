package ViewModels

type UserInputViewModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserOutputViewModel struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
}