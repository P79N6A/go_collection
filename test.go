package main

import "fmt"

// define a struct called Book
type Book struct {
	name string
	price int
}

func testSlice() {

	// init a slice
	slice1:=[]int{}

	// append one element in the tail of the slice
	slice1=append(slice1,4)
	slice1=append(slice1,5)

	// append several elements in the tail of the slice
	slice1=append(slice1,6,7,8)

	// use range to go through the slice
	// first return value is index,second return value is value
	for index,value := range slice1 {
		fmt.Println("index",index,"value",value)
	}
}

func testMap() {

	// init a map
	map1:=map[string]int{}

	// add key-value in the map
	map1["hello"]=1
	map1["hi"]=2

	// use range to go through the map
	// first return value is key
	// second return value is value
	for key,value := range map1 {
		fmt.Println("key is",key,"and value is",value)
	}
}

func testStruct() {

	// init a Book var
	book1 :=Book{"mysql",100}

	// init a Book var
	var book2 Book
	book2.name="sql server"
	book2.price=100

	fmt.Println(book1,book2)

	// test Book's method
	fmt.Println("func call :price is :",book1.call())
}

// this func belongs to struct Book
func (book Book) call() int {
	fmt.Println("this is book:",book.name)
	return book.price
}

func testPointer() {
	
	// Book pointer
	var book3Ptr *Book
	
	book5Ptr := &Book{"gearman",150}
	fmt.Println(book5Ptr)

	book4 := Book{"kafka",120}

	// give book4's address to this pointer 
	book3Ptr = &book4
	fmt.Println(book3Ptr)

	// golang , book3Ptr.call() == *(book3Ptr).call()
	book3Ptr.call()
	
	int1,int2 := 1,2
	var int2Ptr *int
	int2Ptr =&int2

	// print int1's address
	fmt.Println(int1,int2Ptr)
}

// slice just give a copy to the var in this function
// the slice is not influenced
func testSliceCopy(slice []int) {
	sliceTest := slice
	sliceTest=append(sliceTest,1,2,3)
	fmt.Println("test slice is : ",sliceTest)
}

func main() {
	testSlice()
	testMap()
	testStruct()
	testPointer()
    
    
	slice := []int{1,2,3}
	testSliceCopy(slice)
	fmt.Println("after being used ,slice is : ",slice)

	// slice just give a copy to updateSlice,it is not influenced
	updateSlice := slice
	updateSlice=append(updateSlice,4,5,6)
	fmt.Println("update updateSlice",updateSlice)
	fmt.Println("after being given to updateSlice and being updated,slice is : ",slice)

	fmt.Println("hello")
}
