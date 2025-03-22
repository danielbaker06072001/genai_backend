package Logic

import (
	"context"
	"fmt"
	"genai2025/DTO"
	Initializers "genai2025/Initializer"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserLogic(param DTO.UserInputDTO) (*DTO.UserOutputDTO, error) {
	var result DTO.UserOutputDTO

	collection := Initializers.MongoDatabase.Collection("user")

	user := map[string]interface{}{
		"username": param.Username,
		"email":    param.Email,
	}

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Insert result: %+v\n", insertResult)
	insertedID := insertResult.InsertedID
	result.UserId = insertedID.(primitive.ObjectID).Hex()
	result.Username = param.Username
	result.Email = param.Email

	return &result, nil
}