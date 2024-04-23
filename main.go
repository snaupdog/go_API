package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
"github.com/gin-contrib/cors"
)



type FitnessData struct {

	ID     int    `json:"ID"`
	Date          string `json:"date"`
	HeartRate     int    `json:"heart_rate"`
	StepCount     int    `json:"step_count"`
}


var fitnessData []FitnessData

func getData(c *gin.Context){
	c.JSON(http.StatusOK, fitnessData)
}


func addData(c *gin.Context){

		var newData FitnessData
		if err := c.BindJSON(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Add the new data to the slice
    newData.ID = len(fitnessData)+1
    println(newData.ID,newData.Date)
		fitnessData = append(fitnessData, newData)
    c.JSON(http.StatusCreated,newData)
}



func main() {
    gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
 router.Use(cors.Default())
	// router.Use(cors.Default())


    router.POST("/data",addData)
    router.GET("/data",getData)

	router.StaticFile("/", "./index.html")

	router.Run(":8080")
}