package main

import (
    "github.com/gin-gonic/gin"
    "gestion-de-mascotas/internal/db"
    "gestion-de-mascotas/internal/pets"
)

func main() {
    // Initialize database connection
    db.Connect()

    // Create a new Gin router
    router := gin.Default()

    // Set up routes
    router.POST("/pets", pets.AddPet)

    // Start the server
    router.Run(":8080")
}