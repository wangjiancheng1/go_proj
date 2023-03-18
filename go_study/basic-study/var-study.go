package main

// 导入包
import "fmt"

// 全局变量
var c int = 100

/*
*运行启动入口
 */
func main() {
	//定义一个变量
	var name string = "Hello world"
	name = "Hi"
	var age int = 18

	// 定义多个变量
	var (
		name1 string
		age1  int
		addr1 string
	)

	// := 自动推导类型
	name2 := "dad"
	age2 := 13

	// 打印一句话，换行
	fmt.Println(name, age)
	fmt.Println(name1, age1, addr1)
	fmt.Println(name2, age2)
	// 打印类型
	fmt.Printf("%T,%T", name2, age2)
	// 打印内存地址 &取地址 %p内存地址
	fmt.Printf("num:%d,内存地址:%p \n", age2, &age2)

	//变量交换
	var a int = 100
	var b int = 200
	b, a = a, b //快速交换
	fmt.Println(a, b)

	//局部变量
	var c int = 200
	fmt.Println(c)

	// 常量 const 通常大写
	const URL string = "www.baidu.com" //显示定义
	const URL2 = "www.douyin.com"      // 隐

	const c1, c2, c3 = "1", 2, false

	// iota 常量计数器 对常量进行计数
	const (
		i1 = iota
		i2 = iota
		i3 = 100
		i4 = 200
		i5 = iota
	)
	fmt.Println(i1, i2, i3, i4, i5)
}
