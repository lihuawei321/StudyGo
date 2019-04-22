package main

import "fmt"

type person7 struct {
	id int
	name string
	age int
	sex string

}

type student7 struct {
	person7

	class int
	score int
	addr string
}

func (p *person7)Print()  {
	fmt.Printf("编号是：%d\n",p.id)
	fmt.Printf("姓名是：%s\n",p.name)
	fmt.Printf("年龄是：%d\n",p.age)
	fmt.Printf("性别是：%s\n",p.sex)


}


func main() {

	p:=person7{1001,"孙尚香",32,"女"}

	p.Print()
	s:=student7{person7{1002,"糜夫人",32,"女"},3001,100,"巴蜀"}

	//子类继承于父类 可以继承结构体的成员(属性)  也可以继承父类的方法
	//父类不能继承于子类
	s.Print()

}
