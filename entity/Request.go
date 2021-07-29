package entity

// swagger:parameters getBookById updateBook deleteBookById
type IDPath struct {

	// 需要查询/修改的id
	//
	// in: path
	Id int
}

// swagger:parameters addBook
type BookAORequest struct {
	// in: body
	//
	BookAO BookAO
}
type BookAO struct {
	// 书籍名称
	//
	// Required: true
	// Maximum length: 20
	// example: Spring in Action
	Name   string  `json:"name" binding:"required,max=20"`
	// 书籍价格
	//
	// Required: true
	// example: 88.88
	Price  float64 `json:"price" binding:"required,number"`
	// 书籍作者
	//
	// Required: true
	// Maximum length: 12
	// Minimum length: 6
	// example: Craig Walls
	Author string  `json:"author" binding:"required,min=6,max=12"`
}



// swagger:parameters updateBook
type BookUORequest struct {

	// 需要修改的id
	//
	// Required: true
	// in: path
	// example: 1
	id uint

	// in: body
	//
	BookUO BookUO
}



type BookUO struct {
	// 书籍名称
	//
	// Maximum length: 20
	// example: Spring in Action
	Name   string  `json:"name" binding:"max=20"`
	// 书籍价格
	//
	// example: 88.88
	Price  float64 `json:"price" binding:"number"`
	// 书籍作者
	//
	// Maximum length: 12
	// Minimum length: 6
	// example: Craig Walls
	Author string  `json:"author" binding:"min=6,max=12"`
}