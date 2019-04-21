package main

import "fmt"

//func swap(a, b int) {
//	a, b := b, a
//}

//指针变量作为函数参数
func swap(a *int, b *int) {
	fmt.Println(*a)
	fmt.Println(*b)

	//&变量 取地址操作  引用操作
	//*指针变量 取值操作 解引用运算符
	temp := *a
	*a = *b
	*b = temp
}

func main() {

	//a := 10.9 //float64
	//通过自动推导类型创建指针变量
	//所有的指针类型都存储的是一个无符号十六进制整型数据
	//p := &a //*float64
	//*int类型
	//fmt.Printf("%T\n", p)
	//fmt.Println(p)

	a := 10
	b := 20
	//值传递
	//swap(a,b)

	//地址传递 形参可以改变实参的值
	swap(&a, &b)

	fmt.Println(a, b)
}
