package main

import "fmt"

func main0301() {

	fmt.Println("区块链")
	defer fmt.Println("比特币")
	defer fmt.Println("百万")
	fmt.Println("财富自由")
}

func main0302()  {
	a:=10
	b:=20

	//func(a int,b int){
	//	fmt.Println(a+b)
	//}(a,b)
	f:=func(a int,b int){
		fmt.Println(a+b)
	}
	//如果在defer调用时将函数参数先放在内存中 是一个独立的存储空间 不会因为改变值影响数据
	defer f(a,b)

	a=123
	b=321


}

func main()  {
	a:=10
	b:=20

	f:=func( ){
		fmt.Println(a+b)
	}
	//如果不传递参数 使用外部变量 如果外部变量修改 会影响函数的值
	defer f()

	a=123
	b=321
}
