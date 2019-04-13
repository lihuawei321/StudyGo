package main

import "fmt"

func main0301() {

	//switch 变量 {
	//case 值1:
	//	代码体
	//case 值2:
	//	代码体
	//default:
	//	代码体
	//}

	var score int
	fmt.Scan(&score)
	switch score / 10 {

	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}

}

func main0302(){
	var score int
	fmt.Scan(&score)
	switch score > 60{
	case true:
		fmt.Println("及格")
	case false:
		fmt.Println("不及格")

	}
}

func main()  {
	//根据输入的年份月份 计算这个月有多少天

	var y int
	var m int
	fmt.Scan(&y,&m)
	switch m {
	case 1,3,5,7,8,10,12:
		fmt.Println(31)
	case 4,6,9,11:
		fmt.Println("30")
	//fallthrough 在case中向下执行下一个case
	//case 1:
	//	fallthrough
	case 2:
		//判断闰年 能被4整除 但是不能被100整除 或能被400整除
		if (y%4 == 0 && y%100 != 0) ||(y%400 == 0){
			fmt.Println(29)
		}else {
			fmt.Println(28)
		}

	default:
		fmt.Println("erro")

	}
}
