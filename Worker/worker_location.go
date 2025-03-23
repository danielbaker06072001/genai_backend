package Worker

import (
	"context"
	"fmt"
	"genai2025/DTO"
	Initializers "genai2025/Initializer"
	"genai2025/Utils"
	"log"
	"math"
)

type ProximityJob struct {
	Locations []DTO.Location
	Username  string
	Callback  func(*DTO.PromixityJob, error)
}

var JobQueue chan ProximityJob

func InitJobQueue(buffer int) {
	JobQueue = make(chan ProximityJob, buffer)
	go StartWorker(JobQueue)
}

func StartWorker(jobs <-chan ProximityJob) {
	for job := range jobs {
		fmt.Println(job)
		var baseUser *DTO.Location
		for _, loc := range job.Locations {
			if loc.Username == job.Username {
				baseUser = &loc
				break
			}
		}
		if baseUser == nil {
			job.Callback(nil, fmt.Errorf("user not found"))
			continue
		}

		var nearby []DTO.Location
		for _, loc := range job.Locations {
			if loc.Username == baseUser.Username {
				continue
			}
			dist := haversine(baseUser.Latitude, baseUser.Longitude, loc.Latitude, loc.Longitude)
			if dist <= 10.0 {
				nearby = append(nearby, loc)
			}
		}

		fmt.Printf("Nearby locations for user %s: %+v\n", job.Username, nearby)
		job.Callback(&DTO.PromixityJob{UserData: nearby}, nil)

		// ? After finished retrieving close devices, use AI to analyze the data
		// ? Then, return the result to the callback function 
		// ? Input: @username
		go func(username string, nearbyCopy []DTO.Location) { 
			fmt.Printf("AI analysis for user %s\n", username)
			// ! Perform AI Analysis here - Geting the User's profile and close devices
			baseUser, err := GetUserProfile(username)
			if err != nil {
				log.Printf("Error retrieving user profile for %s: %v\n", username, err)
				return
			}
			
			var profiles []DTO.UserPromptDTO
			for _, loc := range nearbyCopy {
				profile, err := GetUserProfile(loc.Username)
				if err != nil {
					log.Printf("Error retrieving user profile for %s: %v\n", loc.Username, err)
					continue
				}
				profiles = append(profiles, *profile)
			}
			result := <- Utils.RankCloseDevicesAsync(*baseUser, profiles)
			if result.Error != nil {
				log.Println("Error:", result.Error)
			} else {
				log.Println("Cohere Response:", result.Response.Text)
			}	
		}(job.Username, nearby)
	}
}


func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func GetUserProfile(username string) (*DTO.UserPromptDTO, error) {
	collection := Initializers.MongoDatabase.Collection("user_profile")

	var profile DTO.UserPromptDTO
	err := collection.FindOne(context.Background(), map[string]interface{}{"username": username}).Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile for user %s: %v", username, err)
	}

	return &profile, nil
}
