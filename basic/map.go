package main

import "fmt"

func TestMapInit() map[string]int {

    //define a map
	map1 := map[string]int{}

	//put k-v in map
	map1["r1"] = 1
	map1["r2"] = 2
	map1["r3"] = 3

	//go through the map
	for key,value := range map1 {
	    fmt.Println("key is ",key," value is ",value)
	}

	return map1
}

func CheckIfKeyExists(map1 map[string]int) {
    value, ok := map1["r1"]
    if(ok) {
    	fmt.Println("r1's value is ",value)
    } else {
        fmt.Println("r1 not exists")
    }
}

func DeleteKV(map1 map[string]int) {
	//传递的是指针，如果删除map2的k-v，map1也会受影响
	map2 := map1

    fmt.Println("删除前：",map2)
    delete(map1,"r1")
	fmt.Println("删除后：",map2)
}

func main() {
	testMap := TestMapInit()
    DeleteKV(testMap)
	CheckIfKeyExists(testMap)
}
