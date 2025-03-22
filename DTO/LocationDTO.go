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

type Location struct {
	Username  string
	Latitude  float64
	Longitude float64
}

type PromixityJob struct {
	UserData []Location
}
