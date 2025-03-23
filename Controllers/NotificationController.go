package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type TokenRequest struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type NotificationRequest struct { 
	Sender string `json:"sender"`
	Receivers []string `json:"receivers"`
	Message string `json:"message"`
}

var userPushTokens = make(map[string]string)

func SendNotification(c *gin.Context) {
	var req NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	success := 0
	failed := 0

	for _, receiver := range req.Receivers {
		payload := map[string]interface{}{
			"subID":    receiver,
			"appId":    28566,
			"appToken": "CxKTyFzipAqvpDDOWwZMBA",
			"title":    fmt.Sprintf("Message from %s", req.Sender),
			"message":  req.Message,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post("https://app.nativenotify.com/api/notification")

		if err != nil || resp.StatusCode() != http.StatusCreated {
			fmt.Printf("Failed to send to user %s\n", receiver)
			failed++
		} else {
			fmt.Printf("Successfully sent to user %s\n", receiver)
			success++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "done",
		"sent":    success,
		"failed":  failed,
		"total":   len(req.Receivers),
		"from":    req.Sender,
		"message": req.Message,
	})
}
