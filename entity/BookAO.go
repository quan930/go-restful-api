package entity

type BookAO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name" binding:"required,max=20"`
	Price  float64 `json:"price" binding:"required,number"`
	Author string  `json:"author" binding:"required,min=6,max=12"`
}
