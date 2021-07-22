package dao

import (
	"bookapi/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type BookDAOImpl struct {

}

var db *sql.DB
func init() {
	fmt.Println("数据库初始化")
	//1、打开数据库
	//parseTime:时间格式转换(查询结果为时间时，是否自动解析为时间);
	// loc=Local：MySQL的时区设置
	sqlStr := "root:quan@tcp(lilq.cn:3306)/golangBook?charset=utf8&parseTime=true&loc=Local"
	var err error
	db, err = sql.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库打开出现了问题：", err)
		return
	}
	//2、 测试与数据库建立的连接（校验连接是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		return
	}
}

func (receiver BookDAOImpl) SelectBooksAll() *[]entity.Book {
	rows, err := db.Query("select id,name,author,price from book")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	books:=make([]entity.Book,0)
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID,&book.Name,&book.Author,&book.Price); err != nil {
			fmt.Println(err)
			return nil
		}
		books = append(books,book)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return &books
}
func (receiver BookDAOImpl) SelectBookById(id string) *entity.Book {
	sqlStr := "select id,name,author,price from book where id=?"
	var book entity.Book
	err := db.QueryRow(sqlStr, id).Scan(&book.ID,&book.Name,&book.Author,&book.Price)
	switch {
		case err == sql.ErrNoRows:
			fmt.Println("No user with that ID.")
		case err != nil:
			fmt.Println(err)
		default:
			return &book
	}
	return nil
}
func (receiver BookDAOImpl) InsertBook(book entity.Book) *entity.Book {
	sqlStr := "insert into book(id,name,author,price) values (?,?,?,?)"
	result, err := db.Exec(sqlStr, book.ID, book.Name, book.Author, book.Price)
	if err != nil {
		fmt.Println("数据库异常")
		return nil
	}
	count, err := result.RowsAffected()
	if count==0 {
		fmt.Println("插入异常")
		return nil
	}else {
		return &book
	}
}
func (receiver BookDAOImpl) UpdateBookById(id string,book entity.Book) *entity.Book {
	//todo 动态sql
	sqlStr := "update book set name=?,author=?,price=? where id = ?"
	result, err := db.Exec(sqlStr, book.Name, book.Author, book.Price,id)
	if err != nil {
		fmt.Println("数据库异常")
		return nil
	}
	count, err := result.RowsAffected()
	if count==0 {
		fmt.Println("修改异常")
		return nil
	}else {
		book.ID = id
		return &book
	}
}
func (receiver BookDAOImpl) DeleteBookById(id string) *int64 {
	sqlStr := "delete from book where id = ?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("数据库异常")
		return nil
	}
	count, err := result.RowsAffected()
	if count==0 {
		fmt.Println("删除异常")
		return nil
	}else {
		return &count
	}
	return nil
}