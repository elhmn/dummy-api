package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Read read data from the data storage
func Read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I read something!"})
}
