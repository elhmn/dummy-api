package handlers

import (
	"net/http"
	"strconv"
	"time"

	"dummy-api/internal/stats"

	"github.com/gin-gonic/gin"
)

//Read read data from the data storage
func Read(c *gin.Context) {
	stats.IncrementHTTPServerReadRequest(c)

	//get latency from the url
	latencyQuery := c.DefaultQuery("latency", "0")
	latency, err := strconv.Atoi(latencyQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//get the `shouldFail` query from the url
	//if shouldFail is true, then the request will fail
	shouldFailQuery := c.DefaultQuery("shouldFail", "false")
	shouldFail, err := strconv.ParseBool(shouldFailQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if shouldFail {
		stats.IncrementHTTPServerErrorRead(c)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail"})
		return
	}

	//wait for the latency
	time.Sleep(time.Duration(latency) * time.Millisecond)
	stats.IncrementHTTPServerOKRead(c)
	c.JSON(http.StatusOK, gin.H{"message": "I read something!"})
}
