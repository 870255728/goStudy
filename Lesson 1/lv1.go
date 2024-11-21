package main

import (
	"fmt"
)

func main() {
	// 定义图形的宽度和高度
	width, height := 400, 200

	for y := height / 2; y >= -height/2; y-- {
		for x := -width / 2; x <= width/2; x++ {
			// 使用心形的方程来确定是否打印 *
			if (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y <= 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
