package main

import (
	"gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize router
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start the HTTP server
    log.Println("Server started at :8080")
    log.Fatal(router.Run(":8080"))
}