package main

import "fmt"

func main() {
	// 算数运算符
	var a int = 128
	var b int = 64
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)
	// 关系运算符 bool true false
	fmt.Println("**********")
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	// 逻辑运算符
	fmt.Println("**********")
	var d bool = false
	var e bool = true
	fmt.Println(d && e)
	fmt.Println(d || e)
	fmt.Println(!e)
	// 位运算符
	fmt.Println("**********")
	fmt.Printf("%d,%b\n", a, a)
	fmt.Printf("%d,%b\n", b, b)
	fmt.Printf("%d,%b\n", b&a, b&a)
	fmt.Printf("%d,%b\n", b|a, b|a)
	fmt.Printf("%d,%b\n", b^a, b^a)
	fmt.Printf("%d,%b\n", a<<2, a<<2)
	fmt.Printf("%d,%b\n", b>>2, b>>2)
	m := -4
	fmt.Printf("%b", m)

	// 赋值运算符
	var n int = 1
	fmt.Println(n)
	// 其他运算符
	var p *int = &n
	fmt.Printf("%p\n", p)
	fmt.Printf("%d\n", *p)
}
