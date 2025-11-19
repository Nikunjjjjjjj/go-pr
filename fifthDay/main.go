package main

//closure,struct(area),method and pointer indirection

// & (address-of) → gets a pointer to a value
// * (dereference) → accesses the value that a pointer refers tov
import (
	"fmt"
	"strings"
)

// Function returning a closure
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type rectangle struct {
	length int
	breath int
}

func main() {

	// counter1 := makeCounter()
	// counter2 := makeCounter()

	// fmt.Println(counter1()) // 1
	// fmt.Println(counter1()) // 2
	// fmt.Println(counter2()) // 1 (separate state)

	// x := 5
	// fmt.Println(x)
	// defer func() {
	// 	fmt.Println(x)
	// }()
	// x = 10
	// fmt.Println(x)
	//wordcount("my name is nikunjj ")

	//s()
	rec1 := rectangle{8, 2}
	// fmt.Println(rec1.area())

	// fmt.Println(rec1.perimeter())

	//fmt.Println(rec1.scale(2))
	fmt.Println(rec1)
	rec1.scale(5)
	fmt.Println(rec1)

}

// When you create a closure:

// Go allocates the captured variables on the heap instead of the stack.

// The closure function keeps a reference (not a copy) of those variables.

// Even after the original function returns, the captured variables stay accessible.

// That’s how Go “remembers” context safely.
// func s() {
// 	// type rectangle struct {
// 	// 	length int
// 	// 	breath int
// 	// }
// 	rec1 := rectangle{1, 2}
// }

func (rec rectangle) area() int {
	return rec.length * rec.breath
}

func (rec rectangle) perimeter() int {
	return 2 * (rec.breath + rec.length)
}

func (rec *rectangle) scale(x int) {
	rec.length = rec.length * x
	rec.breath = rec.breath * x
	// fmt.Println(rec.length, rec.breath)
	// fmt.Printf("%+v", rec)
}

func wordcount(str string) {
	words := strings.Fields(str)
	fmt.Println(words)
	m := make(map[string]int)
	for _, val := range words {
		m[val]++
	}
	fmt.Println(m)
}
