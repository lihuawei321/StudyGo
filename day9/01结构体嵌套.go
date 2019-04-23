package main

import (
	"fmt"
	"unsafe"
)

//结构体不能嵌套自己本身
//可以嵌套结构体本身的指针
//链表的数据存储格式

//线性存储 链式存储  树式存储  图式存储
type student struct {
	*student

	name string
	age int
	sex string
}
func main() {

	var stu student

	//计算数据类型在内存中占的字节大小
	fmt.Println(unsafe.Sizeof(stu))
	//stu.name="张三"
	//
	//fmt.Println(stu)
}
