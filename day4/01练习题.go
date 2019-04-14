package main

import "fmt"

func main() {
	//打印等腰三角形
	//     *        5 6-i-1   1  i*2+1
	//    ***		4	  	  3
	//   *****		3		  5
	//  *******		2	      7
	// *********	1	      9
	//*********** 	0         11

	//行数
	l := 10
	//整体执行次数
	for i := 0; i < l; i++ {
		//控制空格个数
		for j := 0; j < l-i-1; j++ {
			fmt.Print(" ")
		}
		//控制*个数
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
