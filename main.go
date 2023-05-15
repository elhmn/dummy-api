package main

import (
	"context"
	"dummy-api/internal/handlers"
	"dummy-api/internal/stats"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
	}))

	//Add stats middleware to the router
	router.Use(func(c *gin.Context) {
		//add request count stat
		stats.IncrementHTTPServerRequest(c)
		c.Next()
	})

	//Health check endpoint
	router.GET("/health", handlers.Health)
	router.GET("/read", handlers.Read)
	router.POST("/write", handlers.Write)

	ctx := context.Background()
	mp, err := stats.SetupMetrics(ctx, "dummy-api")
	if err != nil {
		panic(err)
	}
	defer mp.Shutdown(ctx)

	if err := router.Run(":7000"); err != nil {
		return
	}
}
