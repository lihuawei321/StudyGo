package main

import "fmt"

func main() {

	//if 表达式
	//{
	//代码体
	//}

	var score int

	fmt.Scan(&score)

	if score > 700 {
		fmt.Println("我要上清华")
	}else if score > 680{
		fmt.Println("我要上北大")
	}else {
		fmt.Println("我要上正大")
	}


}
