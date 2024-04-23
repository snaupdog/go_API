package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type FitnessData struct {
    ID          int    `json:"ID"`
    HeartRate   int    `json:"heart_rate"`
    SleepTime   int    `json:"sleep_time"`
    StepCount   int    `json:"step_count"`
}

var fitnessData []FitnessData

func getData(c *gin.Context) {
    c.JSON(http.StatusOK, fitnessData)
}

func addData(c *gin.Context) {
    var newData FitnessData
    if err := c.BindJSON(&newData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newData.ID = len(fitnessData) + 1
    fitnessData = append(fitnessData, newData)
    c.JSON(http.StatusCreated, newData)
}

func updateData(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var updatedData FitnessData
    if err := c.BindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for index, data := range fitnessData {
        if data.ID == id {
            updatedData.ID = id
            fitnessData[index] = updatedData
            c.JSON(http.StatusOK, updatedData)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
}

func deleteData(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    for index, data := range fitnessData {
        if data.ID == id {
            fitnessData = append(fitnessData[:index], fitnessData[index+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(cors.Default())

    router.POST("/data", addData)
    router.GET("/data", getData)
    router.PUT("/data/:id", updateData)
    router.DELETE("/data/:id", deleteData)

    router.StaticFile("/", "./index.html")

    router.Run(":8080")
}