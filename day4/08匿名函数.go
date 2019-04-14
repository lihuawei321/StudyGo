package main

import (
	"fmt"
)

func main0801() {
	a := 10
	b := 20

	//在函数内部定义一个匿名函数
	//f函数类型变量 接受函数
	var f func(int, int)
	f = func(a int, b int) {
		fmt.Println(a + b)
	}

	f(a, b)
	// fmt.Println(f)
}

func main0802() {
	//a := 10
	//b := 20
	//int 型
	//v := func(a int, b int) int {
	//	return a + b
	//}(a, b)

	f := func(a int, b int) int {
		return a + b
	}

	v := f(10,20)
	fmt.Printf("%T\n",f)
	fmt.Println(v)
	}

func main0803()  {
	a:=10
	b:=20
	f:= func() int{
		return a+b
	}
	//v:=f(a,b)
	a=100
	b=200

	v:=f()

	fmt.Println(v)
}
