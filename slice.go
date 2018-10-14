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

func TestSliceAppend() {
    fmt.Println("start to test append")
    slice1 := []int{1,2,3}
	
	//slice1's cap is  3,len is 3, the address is  0xc00008c020
    fmt.Println("slice1's cap is ",cap(slice1),"slice1's len is",len(slice1)," the address is ",&slice1[0])
    
	
	slice1 = append(slice1,1)
    //slice1's cap is  6,len is 4, the address is  0xc000078090,两倍扩容，地址改变，说明重新创建了一个slice
    fmt.Println("slice1's cap is ",cap(slice1),"slice1's len is",len(slice1)," the address is ",&slice1[0])
    
    slice1 = append(slice1,1)
    //slice1's cap is  6,len is 5, the address is  0xc000078090，目前容量为6，容量足够，不用扩容，不用重新创建一个slice,
    //所以地址未改变
    fmt.Println("slice1's cap is ",cap(slice1),"slice1's len is",len(slice1)," the address is ",&slice1[0])
    
    //slice1's cap is  12 slice1's len is 8  the address is  0xc00001c120，加了三个元素，一共8个元素，
    //容量不够了，需要扩容，依然两倍扩容，重新创建一个slice
    slice1 = append(slice1,1,2,3)
    fmt.Println("slice1's cap is ",cap(slice1),"slice1's len is",len(slice1)," the address is ",&slice1[0])
    
    /*
    * go slice扩容机制：如果切片的容量小于1024个元素，扩容的时候就翻倍增加容量，
    * 一旦元素个数超过1024个，增加因子就变成1.25，即每次增加原来容量的1/4
    */
}

func TestModifyElements() {
    
    fmt.Println("start to test modify ")
    slice1 := []int{1,2,3}
	slice2 := slice1

	//改变slice2的元素，slice1跟着改变,slice1和slice2的地址是相同的
    slice2[0] = 4
    fmt.Println("slice1: ",slice1,",address is ",&slice1[0])
	fmt.Println("slice2: ",slice2,",address is ",&slice2[0])

	fmt.Println("append elements in slice2")
	slice2 = append(slice2,1)
    
    //slice1:  [4 2 3] ,address is  0xc00008e040
    //slice2:  [4 2 3 1] ,address is  0xc00007a0c0
    fmt.Println("slice1: ",slice1,",address is ",&slice1[0])
	fmt.Println("slice2: ",slice2,",address is ",&slice2[0])
	//因为slice2容量与slice1是相同的，slice1在创建时没有指定容量，因此与长度相同，为3
	//当slice2添加元素时，超过了3，因此二倍扩容，重新创建了一个slice，所以地址不再相同
	slice2[0] = 5
	//此时再修改slice2，也不会对slice1产生影响
	//尽量使用这种方法来扩容，即最初创建时不要用make函数指定cap
	//这样每次达到添加元素都会达到最大容量，都会开辟一块内存把原来的值拷贝过来再append()
    fmt.Println("slice1: ",slice1,",address is ",&slice1[0])
	fmt.Println("slice2: ",slice2,",address is ",&slice2[0])
	
	//安全的添加元素的方法
	fmt.Println("test safe append")
	slice5 := []int{}
	slice6 := slice5
	slice6 = append(slice6,1,2,3)
	slice6[0] = 4
	fmt.Println("slice5: ",slice5)
	fmt.Println("slice6: ",slice6)

	//若没有超过容量
	fmt.Println("if not > slice3's cap")
	slice3 := make([]int,2,3)
	slice4 := slice3
	slice4 = append(slice4,1)
	//slice3:  cap is  3 ,len:  2 [0 0] ,address is  0xc0000121a0
    //slice4:  cap is  3 ,len:  3 [0 0 1] ,address is  0xc0000121a0
    //slice3未跟着slice4增加元素
	fmt.Println("slice3: ","cap is ",cap(slice3),",len: ",len(slice3),slice3,",address is ",&slice3[0])
	fmt.Println("slice4: ","cap is ",cap(slice4),",len: ",len(slice4),slice4,",address is ",&slice4[0])

	slice4[0] = 1
	//此时修改slice4也会对slice3产生影响
	//slice3:  cap is  3 ,len:  2 [1 0] ,address is  0xc0000121a0
    //slice4:  cap is  3 ,len:  3 [1 0 1] ,address is  0xc0000121a0
    fmt.Println("slice3: ","cap is ",cap(slice3),",len: ",len(slice3),slice3,",address is ",&slice3[0])
	fmt.Println("slice4: ","cap is ",cap(slice4),",len: ",len(slice4),slice4,",address is ",&slice4[0])
}

func main() {
	slice := TestSliceInit()
    TestSliceIfChange(slice)

    //将slice在函数间传递是值传递，产生了一个新的slice
    //但是新的slice仍然指向原来的底层数组，所以通过新的slice也能改变原来slice的值
    //但是对新的slice做增加元素操作不会增加原来slice的元素
    fmt.Println("slice : ",slice)
    
    fmt.Println("----------")
    TestSliceAppend()
    
    fmt.Println("----------")
    TestModifyElements()
}
