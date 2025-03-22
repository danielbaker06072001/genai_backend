package Logic

import (
	"context"
	"fmt"
	"genai2025/DTO"
	Initializers "genai2025/Initializer"
	"genai2025/Worker"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveLocationLogic(param DTO.LocationInputDTO) (*DTO.LocationOutputDTO, error) {
	var result DTO.LocationOutputDTO

	collection := Initializers.MongoDatabase.Collection("Location")

	location := map[string]interface{}{
		"username":        param.Username,
		"longtitude": param.Longitude,
		"latitude":     param.Latitude,
	}

	// Ensure Uniqueness

	// Check if the username already exists
	existingLocation := collection.FindOne(context.Background(), map[string]interface{}{
		"username": param.Username,
	})
	if existingLocation.Err() == nil {
		_, err := collection.UpdateOne(
			context.Background(),
			map[string]interface{}{"username": param.Username},
			map[string]interface{}{
				"$set": map[string]interface{}{
					"longtitude": param.Longitude,
					"latitude":   param.Latitude,
				},
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to update location for username %s: %v", param.Username, err)
		}

		result.LocationId = "" // No new ID since it's an update
		result.Username = param.Username
		result.Longitude = param.Longitude
		result.Latitude = param.Latitude

		return &result, nil
	}

	// Insert the new location
	insertLocation, err := collection.InsertOne(context.Background(), location)
	if err != nil {
		return nil, err
	}

	insertedID := insertLocation.InsertedID
	result.LocationId = insertedID.(primitive.ObjectID).Hex()
	result.Username = param.Username
	result.Longitude = param.Longitude
	result.Latitude = param.Latitude

	return &result, nil
}

func GetClosestLocationLogic(username string) error {
	collection := Initializers.MongoDatabase.Collection("Location")

	// Get all locations from DB
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to get locations: %v", err)
	}
	defer cursor.Close(context.Background())

	var locations []DTO.Location
	for cursor.Next(context.Background()) {
		var loc DTO.Location
		if err := cursor.Decode(&loc); err != nil {
			continue
		}
		locations = append(locations, loc)
	}

	// ✅ Enqueue the job with username and callback
	job := Worker.ProximityJob{
		Locations: locations,
		Username:  username,
		Callback: func(result *DTO.PromixityJob, err error) {
			if err != nil {
				fmt.Printf("Proximity job failed for %s: %v\n", username, err)
				return
			}

			// ✅ This is where you handle the result
			fmt.Printf("Nearby users for %s: %+v\n", username, result.UserData)

			// Optional: Save result to DB, notify user, or cache it
		},
	}

	Worker.JobQueue <- job
	return nil // you could return 202 accepted or just success
}
