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

func main() {
	time.Sleep(1000)
	//时
	for i := 0; i < 24; i++ {
		//分
		for j := 0; j < 60; j++ {
			//秒
			for k := 0; k < 60; k++ {
				time.Sleep(time.Microsecond * 950)
				fmt.Printf("%d 时 %d 分 &d 秒\n", i, j, k)
			}
		}
	}

}
