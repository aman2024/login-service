package routes

import (
	"login-service/handler"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin" // gin-swagger middleware
)

func InitRoutes(router *gin.Engine) {
	// Cart Specific Routes

	// To be changed to allow only specific Origins

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("session_id")
	router.Use(cors.New(config))
	router.Use(cors.Default())

	// responseBodyTimeout := gin.H{
	// 	"code":    http.StatusRequestTimeout,
	// 	"message": "request timeout, response is sent from middleware"}
	// router.Use(timeout.TimeoutHandler(35*time.Second, http.StatusRequestTimeout, responseBodyTimeout))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "code": http.StatusNotFound, "message": "Link not found or URL not found"})
	})
	api := router.Group("/api/v1/")

	api.GET("/health", handler.GetHealth)

}
