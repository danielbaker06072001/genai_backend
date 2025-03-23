package Controllers

import (
	"genai2025/DTO"
	"genai2025/Logic"
	"genai2025/ViewModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
	})
}


func CreateUser(c *gin.Context) { 
	var userVM *DTO.UserInputDTO
	var response ViewModels.CommonViewModel

	if err := c.ShouldBind(&userVM); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userSponse *DTO.UserOutputDTO
	userSponse , err := Logic.CreateUserLogic(*userVM)
	if err != nil {
		response.Result = "Error"
		response.Message = "Create User"
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Result = "OK"
	response.Message = "Create User"
	response.Data = userSponse

	c.JSON(http.StatusOK, response)

}

func CreateUserProfile(c *gin.Context) {
	var param DTO.UserProfileInputDTO
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	var response ViewModels.CommonViewModel
	userProfileOutput, err := Logic.CreateUserProfile(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response.Result = "OK"
	response.Message = "Create User Profile"
	response.Data = userProfileOutput

	c.JSON(http.StatusOK, response)
}