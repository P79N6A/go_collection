package main

import(
	"fmt"
)

type Car struct {
	speed float64
}

type Benz struct {
	*Car //匿名字段类型
	brand string
}
//Benz结构体声明了匿名字段Car，这个类型的值可以拥有Car的所有方法

func (c *Car) MileToKm() float64 {
	return c.speed*1.6
}

func main() {
	b := Benz{&Car{100},"benz"}
	fmt.Println(b.MileToKm()) //Benz直接调用Car的方法
}

