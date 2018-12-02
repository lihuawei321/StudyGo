package main

import "fmt"

func main0901() {
	a := 10
	b := 20
	c := a + b

	//c += 20// c=c+20
	//c -= 20
	//c *= 20
	//c /= 20
	c = 20
	c %= 3 //c = c % 3

	fmt.Println(c)

}
func main()  {
	var c int = 10
	//将表达式右侧进行结果计算然后在进行赋值运算符
	c %= 2 + 3//c = c % 5
	fmt.Println(c)
}
