package main

import "fmt"

type People struct {
    name string
    age int
    sex string
    job string
}

//struct的构造函数，称之为工厂模式
func NewPeople(Name string,Age int,Sex,Job string) *People {
    return &People{
		name : Name,
        age : Age,
		sex : Sex,
		job : Job,
    }
}

//定义属于People指针的方法
//结构体方法和结构体指针方法区别是一个是值传递，一个是引用传递
//如果不是用指针，在函数内部无法修改struct的值
func (people *People) GetName() string {
	return people.name
}

func (people People) SetName1(name string) {
    people.name = name
}

func (people *People) SetName2(name string) {
    people.name = name
}

//函数的返回值可以有多个
func TestInitStruct() (p1,p2 People,p3 *People)  {
	//1.define a People and init 
	people1 := People{"tom",20,"male","student"}

	//2.define a People then init
	people2 := People{}
    people2.name = "lily"
    people2.age = 18
    people2.sex = "female"

	//3.new a people ptr
	people3 := new(People)
	people3.name = "bill"


	fmt.Println(people1.age,people2.age)
	fmt.Println(people1,people2,people3)

	return people1,people2,people3
}

func TestReceiveTwoReturnType(people1,people2 People) {
	fmt.Println(people1,people2)
}

func TestModifyStruct(people1 People) {
    people2 := people1
    people2.age = 21
    fmt.Println(people2)
}

func TestModifyStructPtr(peoplePtr *People) {
    peopleTestPtr := peoplePtr
    peopleTestPtr.age = 22
    fmt.Println(peoplePtr)
}

func main() {

	//有几个返回值就要用几个变量去接收
	people1,people2,people3 :=  TestInitStruct()
	TestReceiveTwoReturnType(people1,people2)
    fmt.Println(people3)
    TestModifyStruct(people1)
    //传递struct，在函数中修改后不会影响原来的struct
    fmt.Println(people1)
 
 	//传递struct指针，在函数中修改后会影响原来的struct
    TestModifyStructPtr(&people1)
    fmt.Println(people1)

    //用构造函数创建people指针
    people4 := NewPeople("lisa",25,"female","worker")
    fmt.Println(people4)

	//struct和struct指针都可以调用这个方法
	name := people4.GetName()
	name1 := people1.GetName()
	fmt.Println(name,name1)

	//SetName1()是结构体方法，值传递，在函数内部使用不会改变struct的值
    people1.SetName1("corey")
    fmt.Println(people1)

	//SetName2()是结构体指针方法，引用传递，在函数内部使用会改变struct的值
    people1.SetName2("ella")
    fmt.Println(people1)
}
