package entity

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Price  float64 `json:"price"`
	Author string `json:"author"`
}

func (receiver Book) ToString() {
	e, err := json.Marshal(receiver)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}