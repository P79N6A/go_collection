package main

import (
	"fmt"
)

type Car interface {
	Run(name string)
	Stop(name string)
	Wheel
}

type Wheel interface {
	Accelerate()
}

type Michelin struct {
	id string
	price int
}

type Tesla struct {
	id string
	Wheel // 如果我们嵌入了一个类型，方法只需要在所嵌入的类型中实现一次，即可在所有包含该嵌入类型的类型中使用
}

type Benz struct {
	id string
	Wheel
}

func (m *Michelin) Accelerate() {
	fmt.Println(m.id+" is accelerating")
}

func (t *Tesla) Run(name string) {
	fmt.Println(t.id+" is running, made by "+name)
}

func (t *Tesla) Stop(name string) {
	fmt.Println(t.id+" stops, made by "+name)
}

//用interface{}接收任意类型，在接收后检查具体类型是什么
func GetInfo(car ...interface{}) {
	for _,info := range car {
		switch info.(type) {
		case (*Tesla) :
			fmt.Println("tesla ptr")
		case (Benz) :
			fmt.Println("benz")
		default :
			fmt.Println("other type")
		}
	}
}

func main() {
	var c Car
	m := Michelin{id : "michelin_100", price : 1000}
	t := Tesla{id : "tesla_100", Wheel : &m} //用接口去接收实现了它的结构体指针
	c = &t //用接口去接收实现了它的结构体指针
	c.Run("tesla")
	c.Stop("tesla")
	c.Accelerate() //直接用结构体去调用嵌入字段的方法
	var c1 Car
	t1 := Tesla{id : "tesla_101"}
	c1 = &t1
	c1.Run("tesla")
	b := Benz{id : "benz_100", Wheel : &m}
	GetInfo(c,c1,b)
}



