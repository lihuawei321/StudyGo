package userinfo

import "fmt"

//在同级别目录下的包名要相同
//如果采用的是包名.函数名 这种调用结构 需要函数名第一个字母大写
func UserLogin() {
	fmt.Println("用户登录成功")
}
