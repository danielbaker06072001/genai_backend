package Controllers

import (
	"genai2025/DTO"
	"genai2025/Logic"
	"genai2025/ViewModels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveLocation(c *gin.Context) {
	var input *DTO.LocationInputDTO
	var outputVM ViewModels.CommonViewModel
	if err := c.ShouldBind(&input); err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var response *DTO.LocationOutputDTO
	response , err := Logic.SaveLocationLogic(*input)
	if err != nil {
		outputVM.Result = "Error"
		outputVM.Message = "Save Location Not Successfull"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outputVM.Result = "Success"
	outputVM.Message = "Save Location Successfull"
	outputVM.Data = response
	c.JSON(http.StatusOK, outputVM)
}
