package main

import "fmt"

func main() {
	// bool:  true false(默认)
	var isFlag bool
	fmt.Println(isFlag)

	// 数字
	/*
	* byte uint8
	* rune uint32
	* int uint64
	 */
	var byte1 byte = 9 //*
	var rune1 rune = 10
	var int1 int = 10 //*
	fmt.Printf("%T,%T,%T\n", byte1, rune1, int1)

	// float
	var float1 float32 = 23.44
	var float2 float64 = 23.44 //*
	fmt.Printf("%T,%T\n", float1, float2)

	// 字符串 不可修改
	var str1 string = "dasdwadafaw"
	var str2 = 'w' // 单引号 字符 整型int32 ASCII字符码
	fmt.Printf("%s,%T\n", str1, str1)
	fmt.Printf("%d,%T", str2, str2)
	fmt.Println("dada" + "dadad") // 字符串拼接

	fmt.Println(len(str1))
	fmt.Println(str1[0])
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	for i, v := range str1 {
		fmt.Println(i)
		fmt.Println(v)
	}
	// 类型转换 需显示转型
	a := 2
	b := 4.0
	c := float64(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	// 整型不能转化为bool类型
	// 转化后变量 = 转化类型（变量）
}
