package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Write write data to the data storage
func Write(c *gin.Context) {
	//get latency from the url
	latencyQuery := c.DefaultQuery("latency", "0")
	latency, err := strconv.Atoi(latencyQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//wait for the latency
	time.Sleep(time.Duration(latency) * time.Millisecond)

	c.JSON(http.StatusOK, gin.H{"message": "I wrote something!"})
}
