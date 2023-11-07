package database

import (
	"Book/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func InitDB() (*gorm.DB, error) {
	conn := "host=host.docker.internal port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Book{})
	logrus.Info("Successfully connected to database")
	return db, nil
}

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	if err := db.Find(&books).Error; err != nil {
		logrus.Errorf("No book is recorded: %v", err)
		return nil, err
	}
	return books, nil
}

func GetBooksByID(db *gorm.DB, ID int) (*models.Book, error) {
	var book models.Book
	if err := db.First(&book, ID).Error; err != nil {
		logrus.Errorf("No bookID is recorded: %v", err)
		return nil, err
	}
	return &book, nil
}

func SaveBook(db *gorm.DB, book models.Book) error {
	err := db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, book models.Book) error {
	err := db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *gorm.DB, ID int) error {
	var book models.Book
	err := db.First(&book, ID).Error
	if err != nil {
		return err
	}
	return nil
}
