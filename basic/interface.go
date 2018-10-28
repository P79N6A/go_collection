package main

import (
	"fmt"
)

type Phone interface {
	call()
	show()
}

type Nokia struct {
	ptype string
	price int
}

type iPhone struct {
	ptype string
	os string
	price int
}

type Mi struct {
	ptype string
	os string
	price int
}


func (nokia Nokia) call() {
	fmt.Println("nokia call")
}

func (nokia Nokia) show() {
	fmt.Println("this is nokia")
}

func (iphone iPhone) call() {
	fmt.Println("iphone call")
}

func (iphone iPhone) show() {
	fmt.Println("this is iphone")
}

func (iphone iPhone) play() {
	fmt.Println("iphone play")
}

func (mi Mi) call() {
	fmt.Println("mi call")
}

func showPhoneInfo(phone Phone) {
	fmt.Println("==========")
	fmt.Println(phone)
	phone.call()
	phone.show()
	fmt.Println("==========")
}

//func (phone Phone) showBaseInfo() {
//
//}

func main() {
	//定义一个Phone接口类型
	var phone Phone

	//用Phone接口类型去接收Nokia类型，
	//一旦接收了，Nokia类型必须实现Phone的所有方法
	//否则编译会报错
	phone = new(Nokia)
	phone.call()
	phone.show()

	//用Phone接口去接收Nokia类型
	showPhoneInfo(Nokia{"E63",800})

	phone = new(iPhone)
	phone.call()
	phone.show()
	showPhoneInfo(iPhone{"iphone4s","ios5",2100})

	//下面这行代码无法通过编译，因为Mi没有实现接口的方法
	//结构体必须实现接口的所有方法，才可以用接口去接收这个结构体
	showPhoneInfo(Mi{"xiaomi6","android",2800})

	//接口是一种标准，它是一个定义了一些方法的集合
	//任何想要实现这个接口的struct，都必须实现所有这些方法
}

