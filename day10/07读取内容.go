package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main0701() {

	//只读方式打开文件
	fp, err := os.Open("D:/a.txt")

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	defer fp.Close()
	buf := make([]byte, 10)
	//n,_:=fp.Read(buf)
	//块读取
	//fmt.Print(string(buf[:n]))
	//n,_=fp.Read(buf)
	//fmt.Print(string(buf[:n]))
	//n,_=fp.Read(buf)
	//fmt.Print(string(buf[:n]))

	for {
		n, err := fp.Read(buf)
		//io.EOF 表示的是文件的结尾
		//当读取到文件末尾 返回值为errors.New("EOF")
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}

func main()  {

	fp,err:=os.OpenFile("D:/a.txt",os.O_RDONLY,6)
	if err!=nil{
		fmt.Println("文件打开失败")
		return
	}

	defer fp.Close()
	//创建文件缓冲区
	r :=bufio.NewReader(fp)
	//行读取 截取的标志位'\n' 可以使用ASCII码中任意字符
	//slice,_:=r.ReadBytes('\n')
	//
	//fmt.Println(slice)

	for{
		slice,err:=r.ReadBytes('h')

		fmt.Println(string(slice))
		//先打印再判断
		if err == io.EOF{
			break
		}
	}
}
