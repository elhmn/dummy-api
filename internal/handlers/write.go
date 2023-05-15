package handlers

import (
	"dummy-api/internal/stats"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Write write data to the data storage
func Write(c *gin.Context) {
	stats.IncrementHTTPServerWriteRequest(c)

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
		stats.IncrementHTTPServerErrorWrite(c)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail"})
		return
	}

	//wait for the latency
	time.Sleep(time.Duration(latency) * time.Millisecond)
	stats.IncrementHTTPServerOKWrite(c)
	c.JSON(http.StatusOK, gin.H{"message": "I wrote something!"})
}
