package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type EmailJob struct {
	Sender  string `json:"sender" binding:"required"`
	JobId   int    `json:"job_id" binding:"gte=0"`
	Subject string `json:"subject" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, "Email Server Running")
	})
	router.POST("/send-email", func(c *gin.Context) {
		var emailHandler EmailJob
		if err := c.ShouldBindJSON(&emailHandler); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!", "er": err})
			return
		}

		log.Printf("Email Service: creating new Email job from invoice #%v...", emailHandler.Sender)
		rand.Seed(time.Now().UnixNano())
		emailHandler.JobId = rand.Intn(1000)
		log.Printf("PrintService: created print job #%v", emailHandler.JobId)

		//send email here
		// sendEmail(emailHandler.Subject, emailHandler.Body+" has been shipped")
		//
		c.JSON(200, emailHandler)
	})
	router.Run(":9000")
}
