package main

import "fmt"

func TestArrayInit() [5]int {

    //-------------------------------

    //1.define then init
    //define an array,length is 5 
	arr1 := [5]int{}
    //then init
    arr1[0] = 1;

    //-------------------------------

    //2.define and init

	arr2 := [5]int{1,2,3,4,5}

    //-------------------------------
    //3.use index to put data

	arr3 := [5]int{1:5}

	//-------------------------------

    //4.let go calc the length automatically
	arr4 := [...]int{1,2,3,4,5} 

    //-------------------------------

    //5.let go calc the length automatically,
    //then use index to put data
	arr5 := [...]int{1:5}


	//print arrays
    fmt.Println(arr1,arr2,arr3,arr4,arr5)

	//notice: if return an array , the func return type
	//should include the length of array
    return arr2
}

//notice: if use array as parameters
//it should include the length
func TestIfChange (arr [5]int) {
	arrTest := arr
	arrTest[0] = 6

	//change the arrayTest not influences the arr
    fmt.Println("change is ",arrTest)
    fmt.Println("parameter is ",arr)
}

func main() {
	arr := TestArrayInit()	
    TestIfChange(arr)
}
