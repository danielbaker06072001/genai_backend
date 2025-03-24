package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type TokenRequest struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

// type NotificationRequest struct { 
// 	Sender string `json:"sender"`
// 	Receivers []string `json:"receivers"`
// 	Message string `json:"message"`
// }

var userPushTokens = make(map[string]string)
type NotificationRequest struct {
	Receivers []string `json:"receivers" binding:"required"`
	Sender    string   `json:"sender" binding:"required"`
	Message   string   `json:"message" binding:"required"`
	// Optional fields
	Title     string                 `json:"title"`
	PushData  map[string]interface{} `json:"pushData"`
}

// SendNotification handles sending notifications to users
func SendNotification(c *gin.Context) {
	var req NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	success := 0
	failed := 0

	// Process each receiver
	for _, receiver := range req.Receivers {
		// 1. Create a notification title if not provided
		title := req.Title
		if title == "" {
			title = fmt.Sprintf("You Have Matched with %s", req.Sender)
		}

		// 2. Set up the notification payload
		// Note: We're using just the user ID as subID - Native Notify will match all devices
		// where the subID starts with this value (all composite IDs for this user)
		payload := map[string]interface{}{
			"subID":    receiver,
			"appId":    28566,
			"appToken": "CxKTyFzipAqvpDDOWwZMBA",
			"title":    title,
			"body":     req.Message,
			"dateSent": time.Now().Format("2006-01-02"), // ISO date format
		}

		// 3. Add custom data if provided
		if req.PushData != nil {
			payload["pushData"] = req.PushData
		} else {
			// Default push data with sender info
			payload["pushData"] = map[string]interface{}{
				"sender": req.Sender,
				"type":   "match",
			}
		}

		// 4. Send the notification
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(payload).
			Post("https://app.nativenotify.com/api/notification")

		// 5. Handle response
		if err != nil || resp.StatusCode() != http.StatusCreated {
			fmt.Printf("Failed to send to user %s: %v, status: %d\n", 
				receiver, err, resp.StatusCode())
			failed++
		} else {
			fmt.Printf("Successfully sent to user %s\n", receiver)
			success++
		}
	}

	// Return results
	c.JSON(http.StatusOK, gin.H{
		"status":  "done",
		"sent":    success,
		"failed":  failed,
		"total":   len(req.Receivers),
		"from":    req.Sender,
		"message": req.Message,
	})
}

// SendTargetedNotification sends to a specific user on a specific device
func SendTargetedNotification(c *gin.Context) {
	var req struct {
		UserID   string `json:"userId" binding:"required"`
		DeviceID string `json:"deviceId"` // Optional - if provided, sends to specific device
		Title    string `json:"title" binding:"required"`
		Message  string `json:"message" binding:"required"`
		PushData map[string]interface{} `json:"pushData"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	
	// Create the subID based on whether a specific device is targeted
	subID := req.UserID
	if req.DeviceID != "" {
		subID = fmt.Sprintf("%s@%s", req.UserID, req.DeviceID)
	}

	// Set up payload
	payload := map[string]interface{}{
		"subID":    subID,
		"appId":    28566,
		"appToken": "CxKTyFzipAqvpDDOWwZMBA",
		"title":    req.Title,
		"body":     req.Message,
		"dateSent": time.Now().Format("2006-01-02"),
	}

	if req.PushData != nil {
		payload["pushData"] = req.PushData
	}

	// Send notification
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("https://app.nativenotify.com/api/notification")

	if err != nil || resp.StatusCode() != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  fmt.Sprintf("Failed to send notification: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Notification sent successfully",
		"to":      subID,
	})
}