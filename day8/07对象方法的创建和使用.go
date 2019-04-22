package main

import (
	"fmt"
)

type student5 struct {
	name  string
	age   int
	sex   string
	score int
	addr  string
}

//结构体类型可以作为函数类型
func (s student5) Print()  {
	fmt.Printf("姓名：%s\n",s.name)
	fmt.Printf("性别：%s\n",s.sex)
	fmt.Printf("年龄：%d\n",s.age)
	fmt.Printf("成绩：%d\n",s.score)
}

//函数
func Print()  {
	fmt.Print("hello world")
}


func main() {


	//创建对象
	stu:=student5{"贾政",58,"男",80,"贾府"}

	//对象.方法
	stu.Print()

	stu1:=student5{"贾玲",29,"女",99,"北京"}
	stu1.Print()

	//打印错误日志使用的函数
	//print()
	//函数调用
	//Print()

	//对象的方法名可以和函数名重名
	//fmt.Println(stu.Print)
	//fmt.Println(stu1.Print)
	//打印函数在代码区的地址
	fmt.Println(Print)
}
