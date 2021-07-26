package entity

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name  string
	Price  float64
	Author string
}

func (receiver Book) ToString()  {
	e, err := json.Marshal(receiver)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}