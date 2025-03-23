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

type UserProfileInputDTO struct {
	Username         string   `json:"username"`
	Name             string   `json:"name"`
	Age              int      `json:"age"`
	Degree           string   `json:"degree"`
	University       string   `json:"university"`
	Skills           []string `json:"skills"`
	Certifications   []string `json:"certifications"`
	CompanyInterests []string `json:"companyInterests"`
	Email            string   `json:"email"`
	Hobbies          []string `json:"hobbies"`
	JobTitle         string   `json:"jobTitle"`
}

type UserProfileOutputDTO struct {
	Username         string   `json:"username"`
	Name             string   `json:"name"`
	Age              int      `json:"age"`
	Degree           string   `json:"degree"`
	University       string   `json:"university"`
	Skills           []string `json:"skills"`
	Certifications   []string `json:"certifications"`
	CompanyInterests []string `json:"companyInterests"`
	Email            string   `json:"email"`
	Hobbies          []string `json:"hobbies"`
	JobTitle         string   `json:"jobTitle"`
}