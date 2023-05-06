package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Read read data from the data storage
func Read(c *gin.Context) {
	//get latency from the url
	latencyQuery := c.DefaultQuery("latency", "0")
	latency, err := strconv.Atoi(latencyQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//wait for the latency
	time.Sleep(time.Duration(latency) * time.Millisecond)
	c.JSON(http.StatusOK, gin.H{"message": "I read something!"})
}
