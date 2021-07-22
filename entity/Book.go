package entity

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     string  `json:"id"`
	Name  string  `json:"name"`
	Price  float64 `json:"price"`
	Author string  `json:"author"`
}

func (receiver Book) ToString()  {
	e, err := json.Marshal(receiver)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}