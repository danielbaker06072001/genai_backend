package Controllers

import (
	"genai2025/DTO"
	"genai2025/Logic"
	"genai2025/Utils"

	"github.com/gin-gonic/gin"
)

func AnalyzeProfile(c *gin.Context) {
	var prompt *DTO.ProfilePrompt

	// ! Testing the API
	Utils.Test()

	if err := c.ShouldBindJSON(&prompt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	summary , err := Logic.CohereSummary(prompt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"summary": summary})
}