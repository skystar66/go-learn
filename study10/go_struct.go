package main

import (
	"encoding/json"
	"fmt"
)

func main() {


	struct1()

	struct_json()

}

//定义结构体

type Book struct {

	//标题
	title string
	//作者
	author string
	//主题
	subject string
	//书籍ID
	book_id int
}

func struct1() {

	//声明结构体
	var book1 Book
	//访问结构体成员 结构体复制
	book1.author = "xl"
	book1.book_id = 1
	book1.subject = "test"
	book1.title = "book"
	fmt.Println("变量修改前 title:", book1)
	//结构体指针
	struct1_1(&book1)
	fmt.Println("变量修改后 title:", book1)
}

func struct1_1(book_ref *Book) {
	book_ref.title = "update_book"
}

/*单纯的生成json 首字母大写即可*/
type Book_Json_1 struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

/*生成json 制定生成的filed*/
type Book_Json_2 struct {
	Title   string `json:"titleFiled"`
	Author  string `json:"authorFiled"`
	Subject string `json:"subjectFiled"`
	BookId  int    `json:"book_idFiled"`
	Timstr int64 `json:"-"`
}

func struct_json() {
	var bookJson = Book_Json_1{
		Title:   "xl",
		Author:  "xl",
		Subject: "xl-go",
		BookId:  1,
	}
	if result, err := json.Marshal(&bookJson); err == nil {
		fmt.Println("默认生成json内容:" + string(result))
	}

	var bookJsonFiled = Book_Json_2{
		Title:   "xtitle-filed",
		Author:  "xl-filed",
		Subject: "xl-sub-filed",
		BookId:  1,
		Timstr:  0,
	}

	if result, err := json.Marshal(&bookJsonFiled); err == nil {
		fmt.Println("指定生成json filed内容:" + string(result))
	}
}





