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

type UserPromptDTO struct {
	Username string   `json:"username"`
	Skills   []string `json:"skills"`
	Interest []string `json:"interest"`
}