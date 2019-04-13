package main

import (
	"fmt"
	"time"
)

func main0701() {

	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count++
			fmt.Println(i, j)
		}
	}
	fmt.Println(count)
}

//电子时钟
func main0702() {

	//时
	for i := 0; i < 24; i++ {
		//分
		for j := 0; j < 60; j++ {
			//秒
			for k := 0; k < 60; k++ {
				time.Sleep(time.Millisecond * 950)
				fmt.Printf("%d 时 %d 分 %d 秒\n", i, j, k)
			}
		}
	}

}

//func main() {
//	fmt.Println(time.Now())
//	fmt.Println(time.Now().Second())
//	fmt.Println(time.Now().Day())
//	fmt.Println(time.Now().Month())
//	//这一年已经过了1204天了
//	fmt.Println(time.Now().YearDay())
//}

func main() {
	//九九乘法表
	//外行控制行
	for i := 1; i <= 9; i++ {
		//内层控制列
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
