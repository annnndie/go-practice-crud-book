package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Page   uint   `json:"page"`
}
