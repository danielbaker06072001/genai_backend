package Worker

import (
	"fmt"
	"genai2025/DTO"
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
			if dist <= 1.0 {
				nearby = append(nearby, loc)
			}
		}

		job.Callback(&DTO.PromixityJob{UserData: nearby}, nil)
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