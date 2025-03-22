package Worker

import (
	"fmt"
	"genai2025/DTO"
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
		go func() { 
			fmt.Printf("AI analysis for user %s\n", job.Username)
			// ! Perform AI Analysis here
			testUser := DTO.UserPromptDTO { 
				Username: "user1",
				Skills: []string{"Java", "SQL", "React"},
				Interest: []string{"AI", "Backend systems", "ML pipelines", "REST APIs"},
			}
			testProfiles := []DTO.UserPromptDTO {
				{
					Username: "user2",
					Skills: []string{"Python", "GenerateiveAI"},
					Interest: []string{"AI", "Backend systems", "ML pipelines", "REST APIs"},
				},
				{
					Username: "user3",
					Skills: []string{"Java", "SQL", "React", "Next.js", "Supabase", "REST APIs"},
					Interest: []string{"AI", "Backend systems", "ML pipelines", "REST APIs"},
				},
			}
			result := <- Utils.RankCloseDevicesAsync(testUser, testProfiles)
			if result.Error != nil {
				log.Println("Error:", result.Error)
			} else {
				log.Println("Cohere Response:", result.Response.Text)
			}	
		}()
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