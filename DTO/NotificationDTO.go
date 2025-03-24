package DTO

type NotificationInputDTO struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Data  string `json:"data" binding:"required"`
}

type NotificationOutputDTO struct {
}