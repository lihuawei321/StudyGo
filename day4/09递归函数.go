package main

import "fmt"

//递归函数 函数调用函数本身
//死递归
func demo4(n int) {

	//在函数遇到return表示函数的结束
	if n == 0 {
		return
	}

	demo4(n - 1)
	fmt.Println(n)
}

func main0901() {

	demo4(10)
}

//计算一个数的阶乘
//var sum int =1
var s = 1

func demo5(n int) {
	if n == 1 {
		return
	}
	//s *= n//5*4*3*2
	demo5(n - 1)
	s *= n //2*3*4*5
}
func main() {

	a := 5
	demo5(a)
	fmt.Println(s)
}
