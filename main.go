package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type FitnessData struct {
	Date          string `json:"date"`
	HeartTracker  []int  `json:"heart_tracker"`
	CaloriesBurned int    `json:"calories_burned"`
	StepCount     int    `json:"step_count"`
}

var fitnessData = []FitnessData{
	{
		Date:           "2024-04-19",
		HeartTracker:   []int{78, 80, 76, 82},
		CaloriesBurned: 350,
		StepCount:      8000,
	},
	{
		Date:           "2024-04-18",
		HeartTracker:   []int{78, 80, 76, 82},
		CaloriesBurned: 400,
		StepCount:      10000,
	},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.StaticFile("/", "./index.html")

	// Endpoint to retrieve fitness data
	router.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, fitnessData)
	})

	// Endpoint to add fitness data
	router.POST("/addData", func(c *gin.Context) {
		var newData FitnessData
		if err := c.BindJSON(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Add the new data to the slice
		fitnessData = append(fitnessData, newData)
		c.Status(http.StatusOK)
	})

	router.Run(":8080")
}