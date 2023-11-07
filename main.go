package main

import (
	"Book/database"
	"Book/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %s", err)
	}
	defer db.Close()
	r := gin.Default()

	routes.SetUpRouter(r, db)
	// Custom 404 page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})

	// Start the Gin server
	r.Run(":8080")

}
