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
	var userVM DTO.UserInputDTO
	var response ViewModels.CommonViewModel

	if err := c.ShouldBind(&userVM); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	userReponseDTO := Logic.CreateUserLogic(userVM)
	if userReponseDTO.UserId == "" {
		response.Result = "Error"
		response.Message = "Create User"
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Result = "OK"
	response.Message = "Create User"
	response.Data = userVM

	c.JSON(http.StatusOK, response)

}