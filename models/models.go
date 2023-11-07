package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookName string
	Author   string
}
