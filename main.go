package main

import (
	"fmt"
	"login-service/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(router)

	err := router.Run(":" + "8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
