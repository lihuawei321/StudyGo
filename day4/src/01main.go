package main

//通过命令编译程序
//go build 01main.go 02test.go 03demo.go

//编译当前项目下所有文件
//go build ./src

import "userinfo"
//在多个文件进行编程时 在gobuild中选择配置 在配置选择目录(指定到文件所在的目录级别)进行编程
func main() {

	//函数的作用域是项目中整个文件
	test(10, 20)
	demo()

	//如果是其他包下的文件 需要使用包名.函数名
	userinfo.UserLogin()
	userinfo.UserCreat()


}
