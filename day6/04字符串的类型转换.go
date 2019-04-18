package main

import (
	"fmt"
	"strconv"
)

func main0401() {

	str := "hello world"

	//将字符串转成字符切片  强制类型转换
	slice := []byte(str)

	//fmt.Println(slice)

	slice[4] = 'a'
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%c", slice[i])
	}
}

func main0402() {
	//字符切片
	slice := []byte{'h', 'e', 'l', 'l', 'o', 97}

	fmt.Println(slice)

	fmt.Println(string(slice))
}

func main0403() {
	//将其他类型转成字符串  Format
	//b:=true
	//str:=strconv.FormatBool(b)
	//fmt.Println(str)

	//str := strconv.FormatInt(123, 10)//进制在计算机中进制可以表示2-36进制 2 8 10 16
	//fmt.Println(str)

	//str :=strconv.FormatFloat(3.14159,'f',4,64)
	//fmt.Println(str)

	//iota itoa
	//str:=strconv.Itoa(123)//将十进制转成字符串
	//fmt.Println(str)
}

func main() {
	//将字符串转成其他类型
	//b, err := strconv.ParseBool("false")
	//if err != nil {
	//	fmt.Println("类型转换出错")
	//} else {
	//	fmt.Println(b)
	//}

	//value,_ :=strconv.ParseInt("123",10,64)
	//fmt.Println(value)

	//value,_ := strconv.ParseFloat("3.14159",64)
	//fmt.Println(value)

	value,_ :=strconv.Atoi("123")
	fmt.Println(value)
}
