package main

//func 函数名(参数列表) (返回值类别){
//代码体
//}

//函数定义
//在整个项目中，函数是唯一的 不能重名
func add(s1 int, s2 int) {
	sum := s1 + s2
	println(sum)
}
func main() {
	//fmt.Println("你好")
	//v:=len("你好")
	//fmt.Println(v)
	a := 10
	b := 20

	//函数调用  函数可以多次调用
	//在函数调用时参数为实际参数（实参）有具体的值 用来给形式参数（形参）传递数据
	add(a, b)
	add(1, 2)
}
