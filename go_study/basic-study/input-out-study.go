package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("请输入两个数值啊a,b:")
	fmt.Scanln(&a, &b)
	fmt.Println("****************")
	fmt.Println(a, b)

}
