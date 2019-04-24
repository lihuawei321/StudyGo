package main

import (
	"fmt"
	"os"
)

func main() {

	//\反斜杠在程序中表示转义字符 会将后面一个字符进行转义 [\\]表示一个[\]
	//在写路径时可以使用/正斜杠来代替\反斜杠
	fp, err := os.Create("D:/a.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	//关闭文件
	defer fp.Close()

	//将字符串写入到文件中
	//\n不会换行 原因是 在windows文本文件中 换行是以\r\n  在linux中换行是以\n
	fp.WriteString("hello world\r\n")
	//不会换行
	fp.WriteString("区块链技术")
}
