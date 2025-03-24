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

func CreateUserProfile(param DTO.UserProfileInputDTO) (*DTO.UserProfileOutputDTO, error) {

	collection := Initializers.MongoDatabase.Collection("user_profile")

	userProfile := map[string]interface{}{
		"username" :        param.Username,
		"name":             param.Name,
		"age":              param.Age,
		"degree":           param.Degree,
		"university":       param.University,
		"skills":           param.Skills,
		"certifications":   param.Certifications,
		"companyInterests": param.CompanyInterests,
		"email":            param.Email,
		"hobbies":          param.Hobbies,
		"jobTitle":         param.JobTitle,
	}

	existingProfile := collection.FindOne(context.Background(), map[string]interface{}{
		"username": param.Username,
	})
	if existingProfile.Err() == nil {
		_, err := collection.UpdateOne(
			context.Background(),
			map[string]interface{}{"username": param.Username},
			map[string]interface{}{
				"$set": map[string]interface{}{
					"username":         param.Username,
					"name":             param.Name,
					"age":              param.Age,
					"degree":           param.Degree,
					"university":       param.University,
					"skills":           param.Skills,
					"certifications":   param.Certifications,
					"companyInterests": param.CompanyInterests,
					"hobbies":          param.Hobbies,
					"jobTitle":         param.JobTitle,
				},
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to update profile for email %s: %v", param.Email, err)
		}

		var result DTO.UserProfileOutputDTO
		result.Username = param.Username
		result.Email = param.Email
		result.Name = param.Name
		result.Age = param.Age
		result.Degree = param.Degree
		result.University = param.University
		result.Skills = param.Skills
		result.Certifications = param.Certifications
		result.CompanyInterests = param.CompanyInterests
		result.Hobbies = param.Hobbies
		result.JobTitle = param.JobTitle

		return &result, nil
	}

	_, err := collection.InsertOne(context.Background(), userProfile)
	if err != nil {
		return nil, err
	}

	var result DTO.UserProfileOutputDTO
	result.Email = param.Email
	result.Name = param.Name
	result.Age = param.Age
	result.Degree = param.Degree
	result.University = param.University
	result.Skills = param.Skills
	result.Certifications = param.Certifications
	result.CompanyInterests = param.CompanyInterests
	result.Hobbies = param.Hobbies
	result.JobTitle = param.JobTitle

	return &result, nil
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
