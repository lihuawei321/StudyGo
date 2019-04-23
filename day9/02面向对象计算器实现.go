package main

import "fmt"

type Operate struct {
	num1 int
	num2 int
}

type Add struct {
	Operate
}

type Sub struct {
	Operate
}

//加法子类的方法
func (a *Add) Result() int {
	return a.num1 + a.num2
}

//减法子类的方法
func (s *Sub) Result() int {
	return s.num1 - s.num2
}

func main() {
	//创建加法对象
	var a Add
	a.num1 = 10
	a.num2 = 20
	v := a.Result()
	fmt.Println(v)

	//创建减法对象
	var s Sub
	s.num1 = 10
	s.num2 = 20
	v := s.Result()
	fmt.Println(v)
}
