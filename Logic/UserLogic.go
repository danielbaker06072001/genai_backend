package Logic

import (
	"context"
	"fmt"
	"genai2025/DTO"
	Initializers "genai2025/Initializer"
)

func CreateUserLogic(param DTO.UserInputDTO) DTO.UserOutputDTO {
	var result DTO.UserOutputDTO

	collection := Initializers.MongoDatabase.Collection("user")

	user := map[string]interface{}{
		"username": param.Username,
		"email":    param.Email,
	}

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return result
	}

	fmt.Printf("Insert result: %+v\n", insertResult)


	return result
}