package main

import "fmt"

func main() {

	//fmt.Println("hello go")
	//fmt.Print("华威真帅")
	//fmt.Print("帅的鸭匹")
	//format
	a := 10
	b := 3.14159
	//%d是一个占位符 表示输出一个整型数据
	//%f是一个占位符 表示输出一个浮点型数据
	//%f默认保留6位小数 因为浮点型数据不是精准的数据 六位是有效的
	//%.2f保留小数位数为2位 会对第三位小数进行四舍五入
	//\n表示一个转义字符 换行
	fmt.Printf("%d %.3f\n", a, b)

	c := "华威真帅"
	//%s是一个占位符 表示输出一个字符串类型
	fmt.Printf("%s",c)
}
