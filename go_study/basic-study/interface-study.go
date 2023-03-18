package main

import "fmt"

type Phone interface {
	call()
}
type Apple struct {
}

func (apple Apple) call() {
	fmt.Println("苹果手机")
}

type Android struct {
}

func (android Android) call() {
	fmt.Println("安卓手机")
}

type Shape interface {
	area() float64
}

type Rect struct {
	width  float64
	height float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.height
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * 3.14
}
func main() {
	var phone Phone
	phone = new(Apple)
	phone.call()
	phone = new(Android)
	phone.call()

	var shape Shape
	shape = Circle{2.0}
	fmt.Println(shape.area())
	shape = Rect{
		width:  10,
		height: 20,
	}
	fmt.Println(shape.area())
}
