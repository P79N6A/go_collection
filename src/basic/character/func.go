package main

import(
  "fmt"
)

//变长参数函数
//vals是一个int类型的slice
func sum(vals ...int) {
  total := 0
  for _,val range vals {
    total += val
  }
  return total
}

//调用
fmt.Println(sum(1,2,3,4))

//等价于
values := []int{1,2,3,4}
fmt.Println(sum(values...))
