package Logic

import (
	"context"
	"fmt"
	"genai2025/DTO"
	Initializers "genai2025/Initializer"

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