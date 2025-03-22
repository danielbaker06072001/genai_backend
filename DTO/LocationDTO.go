package DTO

type LocationInputDTO struct {
	Username  string `json:"username"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type LocationOutputDTO struct {
	LocationId string `json:"locationId"`
	Username   string `json:"username"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
}