package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Price  float64 `json:"price"`
	Author string `json:"author"`
}

func (receiver Book) ToString() string{
	e, _ := json.Marshal(receiver)
	return string(e)
}