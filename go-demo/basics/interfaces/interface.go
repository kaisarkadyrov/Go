package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base, height, a, b, c float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func calculate(s Shape) {
	fmt.Printf("The area is %.2f\n", s.Area())
	fmt.Printf("The perimeter is %.2f\n", s.Perimeter())
}

func largest(shapes []Shape) Shape {
	best := shapes[0]

	for _, v := range shapes {
		if v.Area() > best.Area() {
			best = v
		}
	}
	return best
}

func main() {
	shapes := []Shape{
		Rectangle{width: 60, height: 20},
		Circle{radius: 15},
		Triangle{base: 10, height: 7, a: 5, b: 6, c: 8},
	}

	for _, v := range shapes {
		calculate(v)
	}

	fmt.Println(largest(shapes))

	// Go смотрит: "функция ждёт тип Shape, а Shape требует метод Area() float64". Потом проверяет у Rectangle — есть такой метод? Есть. Значит Rectangle подходит. Всё.
}
