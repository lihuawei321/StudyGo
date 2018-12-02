package main

import "fmt"

func main0401() {

	fmt.Printf("35%%\n")
	a := 123//十进制整型数据
	//%b 占位符 打印一个数据的二进制格式
	//%o 占位符 打印一个数据的八进制格式
	//%x %X 十六进制
	//10-15用a-f表示 %x a-f  %X A-F
	fmt.Printf("%b\n",a)
	fmt.Printf("%o\n",a)
	fmt.Printf("%x\n",a)
	fmt.Printf("%X\n",a)
	//十进制
	fmt.Println(a)
	fmt.Printf("%d\n",a)
}
func main0402()  {
	//十进制数据
	var a int = 10
	//八进制数据 八进制数据是以0开头 最大值为7
	var b int = 010
	//十六进制数据 十六进制数据是以0x开头
	var c int =0xabc
	//二进制 不能再go语言中直接表示
	//var d int = 1001110
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
func main()  {
	//''引起来的只能存储一个字符
	ch := 'a'
	str := "区块链"
	fmt.Printf("%c\n",ch)
	fmt.Printf("%s\n",str)
	fmt.Println(ch)

	//%p 占位符 打印是一个变量对应的内存地址 是以无符号十六进制整型表示
 	fmt.Printf("%p\n",&ch)
	fmt.Printf("%p\n",&str)

	 a := false
	 //%t 占位符 打印bool类型的值
	 fmt.Printf("%t\n",a)
}
