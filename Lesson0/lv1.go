package main

import "fmt"

func main() {
	fmt.Println("input:")
	var a, b int
	var op string
	fmt.Scanln(&a, &op, &b)
	switch op {
	case "+":
		fmt.Println(a, op, b, "=", a+b)
	case "-":
		fmt.Println(a, op, b, "-", a+b)
	case "*":
		fmt.Println(a, op, b, "*", a*b)
	case "/":
		fmt.Println(a, op, b, "/", a/b)

	}

}
