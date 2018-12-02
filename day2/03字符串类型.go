package main

import "fmt"

func main0301() {
	//var a string = "华威真帅"
	//fmt.Println(a)
	//fmt.Printf("%s\n",a)
	//a := "华威真帅"
	//fmt.Println(a)
	//fmt.Printf("%s",a)
	//双引号引起来的称为字符串
	ch := 'a'
	str := "a"//'a''\0'字符串的结束标志

	fmt.Printf("%c\n",ch)
	//%s打印字符串打印到\0之前的内容
	fmt.Printf("%s\n",str)

}
func main0302()  {
	//len 函数 用来计算字符串中字符的个数 不包含\0 返回值为int类型
	//a := "hello"
	//在go语言中一个汉字占3个字符 为了和linux进行统一处理
	a := "区块链"
	var  count int
	count = len(a)
	fmt.Println(count)
}
func main()  {
	str1 := "区块链"
	str2 := "比特币"

	//字符串的连接 +
	str3 := str1 + str2
	fmt.Println(str3)
}
