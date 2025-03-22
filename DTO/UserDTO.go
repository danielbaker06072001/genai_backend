package DTO

type UserInputDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserOutputDTO struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
}