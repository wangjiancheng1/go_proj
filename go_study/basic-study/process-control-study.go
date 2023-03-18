package main

import "fmt"

func main() {
	// 顺序结构

	// 分支结构
	var a int = 1
	if a > 3 {
		fmt.Println(3)
	} else if a < 0 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
	// switch 无值 默认为true
	// case穿透  fallthrough break 终止穿透

	switch a {
	case 0:
		fmt.Println(0)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
		break
	default:
		fmt.Println("e")
	}

	switch {
	case true:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false")
		fallthrough
	default:
		fmt.Println("default")
	}
	// 循环结构
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	w := 0

	for {
		w++
		fmt.Println(w)
		if w > 20 {
			break
		}
	}

	// 退出 break 继续 continue
	
	// example
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
