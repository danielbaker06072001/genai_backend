package main

import (
	"genai2025/Controllers"
	Initializers "genai2025/Initializer"
	"genai2025/Worker"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToMongoDB()
	Worker.InitJobQueue(10)
}
func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())	
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{	"message": "pong"	})
	})

	User := r.Group("/user")
	{ 
		User.GET("/get", Controllers.GetUsers)
		User.POST("/create", Controllers.CreateUser)
	}

	Location := r.Group("/location")
	{ 
		Location.POST("/save", Controllers.SaveLocation)
		Location.POST("/get-closest", Controllers.GetClosestLocation)
	}


	if err := r.Run(":8090"); err != nil {
		log.Println("Failed to start server")
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}