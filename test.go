package main

import "fmt"

type Book struct {
	name string
	price int
}

func testSlice() {
	slice1:=[]int{}
	slice1=append(slice1,4)
	slice1=append(slice1,5)
	for index,value := range slice1 {
		fmt.Println("index",index,"value",value)
	}
}

func testMap() {
	map1:=map[string]int{}
	map1["hello"]=1
	map1["hi"]=2
	for key,value := range map1 {
		fmt.Println("key is",key,"and value is",value)
	}
}

func testStruct() {
	book1 :=Book{"mysql",100}
	var book2 Book
	book2.name="sql server"
	book2.price=100
	fmt.Println(book1,book2)
	fmt.Println("func call :price is :",book1.call())
}

func (book Book) call() int {
	fmt.Println("this is book:",book.name)
	return book.price
}

func testPointer() {
	var book3Ptr *Book
	book4 := Book{"kafka",120}
	book3Ptr = &book4
	fmt.Println(book3Ptr)
	book3Ptr.call()
	int1:=1
	var int2Ptr *int
	int2Ptr =&int1
	fmt.Println(int2Ptr)
}

func main() {
	testSlice()
	testMap()
	testStruct()
	testPointer()
	fmt.Println("hello")
}
