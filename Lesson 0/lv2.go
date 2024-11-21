package main

import "fmt"

func main() {
	var book map[string]string
	book = make(map[string]string)
	book["xyp"] = "x"
	book["wcl"] = "w"
	fmt.Println("请输入账号：")
	var zhanghao string
	fmt.Scanln(&zhanghao)
	fmt.Println("请输入密码：")
	var mima string
	fmt.Scanln(&mima)
	if book[zhanghao] == mima {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
