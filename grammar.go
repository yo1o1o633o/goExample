package main

import "fmt"

func main() {
	// 定义指定类型变量
	var str1 string
	fmt.Println(str1)
	// 定义指定类型变量, 并赋值
	var a string = "Hello World"
	fmt.Println(a)
	// 定义变量可以不指定具体类型, Go自动判断
	var b = "Hello World"
	fmt.Println(b)


}
