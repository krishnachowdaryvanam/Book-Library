package handlers

import (
	"Book/database"
	"Book/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := database.GetBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server"})
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBooksByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		BookId := c.Param("id")

		id, err := strconv.Atoi(BookId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID"})
			return
		}
		bookId, err := database.GetBooksByID(db, id)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Book was found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			}
		}
		c.JSON(http.StatusOK, bookId)
	}
}

func SaveBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.SaveBook(db, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sucess": "Book sucessfuly Saved"})
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := database.UpdateBook(db, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")

		id, err := strconv.Atoi(bookID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}
		err = database.DeleteBook(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}
