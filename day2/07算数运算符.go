package main

import "fmt"

func main0701() {
	a := 10
	b := 5

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	//两个整数相除得到的结果也是整型
	//在除法计算时 除数不能为0
	fmt.Println(a / b)

	//通过命令编译代码
	//go build 源文件1.go 源文件2.go
	//编译并运行程序
	//go run 源文件1.go 源文件2.go
	//异常
	/*
	1、编辑时异常
	2、编译时异常
	3、运行时异常
	 */
}
func main0702()  {
	a := 10
	b := 2
	//取模运算符 除数不能为0
	//取模运算符 不能对浮点型使用
	c := a % b
	fmt.Println(c)
}

func main()  {
	//自增自减运算符
	//可以对浮点型进行自增自减运算
	a := 10
	//const a = 10 常量不能自增自减
	//a = a + 1
	a++//自增，在变量本身加1
	a--//自减，在变量本身减1
	//自增自减不能出现在表达式中
	//a = a++ + a--
	//二义性 在不同操作系统中运算方式不同 结果可能会产生偏差
	//a = a++ * a-- - a--

	//b := a--//err

	fmt.Println(a)
	//fmt.Println(b)
}
