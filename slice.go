package main

import "fmt"

func TestSliceInit () []int {
	//1.define a slice then use append() to add elements
    //init a slice
	slice1 := []int{}
	//append one element in the tail of the slice
	slice1 = append(slice1,1)
	slice1 = append(slice1,2)

	//append several elements in the tail of the slice
	slice1 = append(slice1,3,4,5)

	//print slice
	fmt.Println("slice1 is ",slice1)

	//use range to go through slice
	for index,value := range slice1 {
	    fmt.Println("index: ",index,",value: ",value);
	}
	//----------------------------------------------------------
	//2.init slice and init
	//the difference between slice and array is
	//if the length of array is not certain,use [...]
	//and slice dosen't need
	slice2 := []int{1,2,3,4,5}
    fmt.Println("slice2 is ",slice2)

    return slice1
}

func TestSliceIfChange(slice []int) {
    
    sliceTest := slice
    sliceTest[0] = 6
    sliceTest = append(sliceTest,7,8,9)
	fmt.Println("change is : ",sliceTest)
	fmt.Println("parameter is : ",slice)
}

func main() {
	slice := TestSliceInit()
    TestSliceIfChange(slice)

    //将slice在函数间传递是值传递，产生了一个新的slice
    //但是新的slice仍然指向原来的底层数组，所以通过新的slice也能改变原来slice的值
    //但是对新的slice做增加元素操作不会增加原来slice的元素
    fmt.Println("slice : ",slice)
}
