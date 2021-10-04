package main

import "fmt"

func main() {
	/*
		结构体作为函数的参数使用
	*/
	b1 := book1{
		name:  "GO",
		price: 87.55,
	}
	b2 := book1{
		name:  "JAVA",
		price: 58.05,
	}

	printBook(b1)
	printBook(b2)
}

//结构体
type book1 struct {
	name  string
	price float64
}

func printBook(book book1) {
	fmt.Printf("书的名字是: %s,书的价格是: %.2f\n", book.name, book.price)
}
