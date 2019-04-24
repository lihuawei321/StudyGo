package main

import "fmt"

func test(i int)  {
	var arr [10]int

	arr[i]=123

	fmt.Println(arr)
}

func main() {

	test(1)

	//fmt.Println("hello world1")
	//fmt.Println("hello world2")
	////当程序遇见panic是自动终止
	//panic("hello world3")
	//fmt.Println("hello world3")
	//fmt.Println("hello world4")
	//fmt.Println("hello world5")


}
