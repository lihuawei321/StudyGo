package main

import "fmt"

func main1001() {

	a := 10
	b := 20
	//大于 > 小于 <
	//fmt.Println(a > b)
	//大于等于 小于等于
	//fmt.Println(a <= b)
	//相等于 == 不等于 !=
	fmt.Println(a != b)
}
func main1002()  {
	a := 10
	b := 20

	//比较运算符返回值类型为bool类型
	c := a + 20 > b

	fmt.Printf("%T\n",c)
	fmt.Println(c)
}
func main()  {
	a := 'a'
	b := 'A'

	fmt.Println(a > b)
}
