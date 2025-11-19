// package main

// import (
// 	"time"
// )

// // interface,errors
// // type Shape interface {
// // 	Area() float64
// // 	Perimeter() float64
// // }

// // // Define a struct
// // type Rectangle struct {
// // 	Width, Height float64
// // }

// // //	type Circle struct {
// // //		radius float64
// // //	}
// // type Circle struct {
// // 	Radius float64
// // }

// // // Implement the interface
// // func (r Rectangle) Area() float64 {
// // 	return r.Width * r.Height
// // }

// // func (r Rectangle) Perimeter() float64 {
// // 	return 2 * (r.Width + r.Height)
// // }

// // //	func (c Circle) Area() float64 {
// // //		return math.Pi * c.Radius * c.Radius
// // //	}
// // func (c Circle) Area() float64 {
// // 	return 3.14 * c.Radius * c.Radius
// // }

// //	func (c Circle) Perimeter() float64 {
// //		return 2 * 3.14 * c.Radius
// //	}
// func main() {
// 	// var s Shape          // s is a Shape interface
// 	// s = Rectangle{10, 5} // Rectangle implements Shape

// 	// var s2 Shape
// 	// s2=Circle{4}
// 	//fmt.Println("Area:", s.Area())
// 	// fmt.Println("Perimeter:", s.Perimeter())

// 	// shapes := []Shape{
// 	// 	Rectangle{10, 5},
// 	// 	Circle{7},
// 	// }
// 	// for _, shape := range shapes {
// 	// 	//fmt.Printf("Area of ", shape, shape.Area())
// 	// 	fmt.Printf("%T Area: %.2f\n", shape, shape.Area())
// 	// }

// }

// type shape interface {
// 	Area() int
// 	Perimeter() int
// }

// type myError struct {
// 	time time.Time
// }

package main

import (
	"fmt"
)

func main() {
	sayHello()
	fmt.Println("Main function running...")

}

func sayHello() {
	fmt.Println("Hello from goroutine!")
}
