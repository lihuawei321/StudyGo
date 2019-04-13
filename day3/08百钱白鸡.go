package main

import "fmt"

func main0801() {
	//百钱白鸡
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken += 3 {
				count++
				if cock+hen+chicken == 100 && 5*cock+3*hen+chicken/3 == 100 {
					fmt.Printf("公鸡： %d  母鸡: %d  小鸡: %d\n", cock, hen, chicken)
				}
			}
		}
	}
	fmt.Println("执行次数： ", count)
}

func main() {
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			chicken := 100 - cock - hen
			if chicken%3 == 0 && 5*cock+3*hen+chicken/3 == 100 {
				fmt.Printf("公鸡： %d  母鸡: %d  小鸡: %d\n", cock, hen, chicken)
			}
		}
	}
	fmt.Println("执行次数：", count)
}
