package dao

import (
	"bookapi/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BookDAOImpl struct {

}

var db *gorm.DB

func init() {
	//db = config.DBConfig{}.DbConfig()
	dsn := "root:quan@tcp(lilq.cn:3306)/golangBook?charset=utf8mb4&parseTime=True&loc=Local"
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//// 迁移 schema
	db1.AutoMigrate(&entity.Book{})
	db = db1
}

func (receiver BookDAOImpl) SelectBooksAll() *[]entity.Book {
	var books []entity.Book
	// 获取全部记录
	result := db.Find(&books)
	if result.Error != nil {
		return nil
	}
	return &books
}
func (receiver BookDAOImpl) SelectBookById(id uint) *entity.Book {
	var book *entity.Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil
	}else if result.RowsAffected!=0 {
		return nil
	}else {
		return book
	}
}
func (receiver BookDAOImpl) InsertBook(book entity.Book) *entity.Book {
	result := db.Create(&book)
	if result.RowsAffected<1 {
		return nil
	}
	db.First(&book, book.ID)
	return &book
}
func (receiver BookDAOImpl) UpdateBookById(id uint,book entity.Book) *entity.Book {
	book.ID = id
	bookMap := make(map[string]interface{})
	if len(book.Name)!= 0 {
		bookMap["Name"] = book.Name
	}
	if len(book.Author)!= 0 {
		bookMap["Author"] = book.Author
	}
	if book.Price!=0 {
		bookMap["Price"] = book.Price
	}
	result := db.Model(&book).Updates(bookMap)
	if result.RowsAffected<1 {
		return nil
	}
	db.First(&book, book.ID)
	return &book
}
func (receiver BookDAOImpl) DeleteBookById(id uint) *int64 {
	result := db.Delete(&entity.Book{}, id)
	if result.RowsAffected<1 {
		return nil
	}
	return &result.RowsAffected
}