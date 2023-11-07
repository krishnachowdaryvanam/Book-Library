package routes

import (
	"Book/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetUpRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/books", handlers.GetBooks(db))
	r.Group("/books/:id", handlers.GetBooksByID(db))
	r.PUT("/updatebooks", handlers.UpdateBook(db))
	r.POST("/saveBooks", handlers.SaveBook(db))
	r.DELETE("/deletebook/:id", handlers.DeleteBook(db))
}
