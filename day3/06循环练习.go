package main

import (
	"fmt"
)

func main0601() {

	//敲7 7的倍数 个位为7 十位为7 需要敲桌子 1-100

	for i := 0;i <= 100; i++ {
		if i%7 ==0 || i%10 ==7 || 1/10 ==7{
			fmt.Println("敲桌子")
		}else {
			fmt.Println(i)
		}
	}

}

func main0603() {
	//水仙花数 一个三位数 各个位数的立方和等于这个数本身
	for i := 100; i <= 999;i ++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		if a*a*a + b*b*b + c*c*c == i {
			fmt.Println(i)

		}
	}
}
