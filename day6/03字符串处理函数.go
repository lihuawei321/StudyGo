package main

import (
	"strings"
	"fmt"
)

func main0301() {
	//查找一个字符串在另外一个字符串中是否出现
	str1 := "hello world"
	str2 := "llo"
	//Contains(被查找字符串，查找字符串)
	//一般用于模糊查找
	b := strings.Contains(str1, str2)
	if b {
		fmt.Println("找到了")
	} else {
		fmt.Println("未找到")
	}
}

func main0302() {
	//字符串切片
	slice := []string{"1234", "4637", "2907", "2308"}
	//字符串连接
	str := strings.Join(slice, "-")
	fmt.Println(str)

}

func main0303() {

	str1 := "hello world"
	str2 := "llo"
	//查找一个字符串在另外一个字符串中第一次出现的位置 返回值为int类型是下标 找不到返回值为-1
	i := strings.Index(str1, str2)

	fmt.Println(i)
}

func main0304() {
	str := "区块链比特币"
	//将一个字符串重复n次
	str1 := strings.Repeat(str, 3)
	fmt.Println(str1)
}

func main0305() {
	str := "区块链比特币区块链"
	//字符串替换 用作于屏蔽敏感词汇
	//如果替换次数小于0，表示全部替换
	str1 := strings.Replace(str, "区块链", "blcokchain", 1)

	fmt.Println(str1)
}

func main0306() {

	//str1 := "1300-199-1433"
	str1 :="123456@qq.com"
	//将字符串按照标志位进行切割变成切片
	slice := strings.Split(str1, "@")

	//将切片进行连接
	//str2:=strings.Join(slice,"")
	fmt.Println(slice[0])
	//fmt.Println(str2)
}

func main0307()  {

	str :="====a=u=ok===="
	//去掉字符串头尾内容
	str1:=strings.Trim(str,"=")

	fmt.Printf("%s",str1)
}

func main0308()  {

	str1:="   are you ok   "
	//去掉字符串中的空格转成切片 一般用作于统计单词个数
	slice:=strings.Fields(str1)

	//fmt.Println(slice)
	//for _,v :=range slice {
	//	fmt.Printf("=%s=\n",v)
	//}
	fmt.Println(len(slice))
}

func main()  {
	//查找
	//bool类型 = strings.Contains(被查找字符串,查找字符串)  查找一个字符串在另外一个字符串中是否出现
	//int类型 = strings.Index(被查找字符串,查找字符串)  查找一个字符串在另外一个字符串中第一次出现的位置 返回值为int类型是下标 找不到返回值为-1

	//分割组合
	//string类型=strings.Join(字符串切片,"标志")  字符串连接
	//[]string类型strings.Split(切割的字符串,"标志") 将字符串按照标志位进行切割变成切片

	//重复和替换
	//string类型=strings.Repeat(字符串,次数)  将一个字符串重复n次
	//string类型=strings.Replace(字符串,被替换字符串,替换字符串,次数) 字符串替换 如果替换次数小于0，表示全部替换

	//去掉内容
	//string类型=strings.Trim(字符串,去掉字符串)  去掉字符串头尾内容
	//[]string类型=strings.Fileds(字符串)  去掉字符串中的空格转成切片 一般用作于统计单词个数
}


//查找
