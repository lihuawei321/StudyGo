package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集

	//继承于Human
	Humaner

	sing(string)
}

type Student struct {
	name string
	sex  string
	age  int
}

func (s *Student) sayhi() {
	fmt.Printf("大家好，我是%s，我是%s生，我的年龄是%d岁\n",
		s.name, s.sex, s.age)
}

func (s *Student) sing(name string) {
	fmt.Println("我为大家唱首歌", name)
}

func main0601() {

	//接口类型变量定义
	var h Humaner

	var stu Student = Student{"王菲", "女", 35}

	h = &stu
	h.sayhi()

	//接口类型变量定义
	var p Personer
	p = &stu

	//从Humaner继承来的
	p.sayhi()
	p.sing("传奇")
}

func main0602() {

	//接口类型变量定义
	var h Humaner
	var p Personer
	var stu Student = Student{"王菲", "女", 35}

	h = &stu
	//将一个接口赋值给另外一个接口
	//超集中包含所有子集的方法
	h = p //ok
	h.sayhi()

	//子集不包含超集
	//可以将超集赋值给子集 不能将子集赋值给超集
	//p = h //err

	//p.sayhi()
	//p.sing("红豆")

}
