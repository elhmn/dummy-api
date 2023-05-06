package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Write write data to the data storage
func Write(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I wrote something!"})
}
