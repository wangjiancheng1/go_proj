package main

import "fmt"

/*
自定义函数

	func 函数名(参数1，参数2...) (函数返回值){
		函数体
	}
*/
func add(a, b int) int {
	return a + b
}

// 无参无返回值
func noCnoR() {
	fmt.Println("0个参数0个返回值函数")
}

// 有一个参数的函数
func oneCnoR(msg string) {
	fmt.Println(msg)
}

// 有多个参数的函数
func moreCnoR(a int, msg string) {
	fmt.Println(a, msg)
}

// 有一个返回值的函数
func moreConeR(a int, msg string) int {
	fmt.Println(a, msg)
	return a
}

// 有多个返回值的函数
func moreCmoreR(a int, msg string) (int, string) {
	fmt.Println(a, msg)
	return a, msg
}

// ...可变参数
func vC(len int, nums ...int) {
	fmt.Println(len, "\b个参数的可变参数函数", nums)
}

/*
参数传递
	值传递   int float string bool array
	引用传递   slice map chan
*/

// 值传递
func valueTrans(arr [4]int) {
	arr[0] = 100
}

// 引用传递
func referenceTrans(arr []int) {
	arr[0] = 100
}

// 递归
func recursive(a int) int {
	if a < 1 {
		return a
	}
	return a + recursive(a-1)
}

// 延迟 defer
func df() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
}

// 函数也是一种类型
func f() {

}

func f1(a, b int) {
	fmt.Println(a, b)
}

// 闭包
func outer() func() int {
	var i int = 0
	f := func() int {
		i++
		return i
	}
	return f
}

func main() {
	// 调用 函数名()
	noCnoR()
	oneCnoR("1个参数0个返回值的函数")
	moreCnoR(2, "\b个参数0个返回值的函数")
	_ = moreConeR(2, "\b个参数2个返回值的函数")
	_, _ = moreCmoreR(2, "\b个参数2个返回值的函数")
	vC(6, 1, 2, 3, 4, 5, 6)

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr, "前")
	valueTrans(arr)
	fmt.Println(arr, "后")

	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr2, "前")
	referenceTrans(arr2)
	fmt.Println(arr2, "后")
	fmt.Println(recursive(10))
	df()
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", f1)
	var ff func(int, int)
	ff = f1
	ff(1, 2)
	// 匿名函数
	fff := func(a int) int {
		fmt.Println("匿名函数")
		return a
	}

	m := func(a int, fun func(int2 int) int) int {
		r := fun(a)
		return r
	}(1, fff)

	fmt.Println(m)
	f111 := outer()
	fmt.Println(f111)
	r1 := f111()
	fmt.Println(r1)
	v2 := f111()
	fmt.Println(v2)
	f222 := outer()
	r21 := f222()
	fmt.Println(r21)
	r3 := f111()
	fmt.Println(r3)
	r22 := f222()
	fmt.Println(r22)
}
