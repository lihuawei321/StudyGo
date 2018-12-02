package main

import (
	"fmt"
	"math"
)

func main0501() {
	//var a int = 10
	//常量的定义 一般定义常量使用大写字母
	const MAX int = 10
	//a = a + 10 常量不允许左值赋值
	fmt.Println(MAX)
	b := MAX + 10
	fmt.Println(b)

}
func main0502()  {
	//const MAX := 10 常量不能使用自动推导定义
	const MAX  = "区块链"
	fmt.Println(MAX)
	fmt.Printf("%T\n",MAX)
	//fmt.Printf("%p\n",&MAX) go语言的地址 不允许访问 //err
}
func main()  {
	//计算圆的面积和周长
	//面积 pai*r*r math.Pow(r,2)
	const pai float64= 3.14
	var r float64
	fmt.Scan(&r)
	s := pai*math.Pow(r,2)
	l := 2*pai*r
	fmt.Printf("面积: %.2f",s)
	fmt.Printf("周长: %.2f",l)

}
func main0504()  {
	//数据区 -> 常量区
	//字面常量
	fmt.Println("hello world")
	fmt.Println(3.14)
	a := 10
	fmt.Println(a + 20)
}