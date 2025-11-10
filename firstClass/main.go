package main

import (
	"fmt"
	"math"
)

func main() {
	// pi := math.Pi
	// fmt.Println(pi)

	// //reverse a string
	// str := "nikunj"
	// newStr := ""
	// for i := len(str) - 1; i >= 0; i-- {

	// 	newStr = newStr + string(str[i])
	// }
	// fmt.Println(newStr)

	//reverse num
	// num := 12345678
	// newNum := 0
	// for num > 0 {
	// 	ld := num % 10
	// 	num = num / 10
	// 	newNum = newNum*10 + ld
	// }
	// fmt.Println(newNum)

	//addfunc
	// s := add(2, 3)
	// fmt.Println(s, reflect.TypeOf(s))

	// fmt.Println(swap("nik", "jain"))
	// var c, pc bool
	// var s string
	// var i int
	// var f float64
	// fmt.Println(c, pc, "\n", s, "\n", i, "\n", f)
	//var i1 = 1

	// i3, i4, i7 := 3, 4, 7
	// var i6, i5 int = 1, 2
	// const i9 = 77
	// fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8+1, i9+1)
	// i2 := 6.98
	// i3 := "vgvhb"
	// fmt.Printf("type:%T,value:%v", i1, i1)
	// fmt.Printf("type:%T,value:%v", i2, i2)
	// fmt.Printf("type:%T,value:%v", i3, i3)
	// var i = 5.6
	// fmt.Println(int64(i))
	// var i2 = 55
	// fmt.Println(reflect.TypeOf(strconv.Itoa(i2)))//%T
	res := sum(5)
	fmt.Println(res)
	res1 := sumW(5)
	fmt.Println(res1)
	res2 := sqroot(4)
	fmt.Println(res2)
	grade := grade(99)
	fmt.Println(grade)
}
func sum(a int) int {
	sum := 0
	for i := 0; i < a; i++ {
		sum += i
	}
	return sum
}

// sum by while
func sumW(a int) int {
	i := 0
	sum := 0
	for i < a {
		sum += i
		i += 1
	}
	return sum
}

// math sq root math.sqroot
func sqroot(a float64) float64 {
	res := math.Sqrt(a)
	return res
}

// grade
func grade(a int) string {
	if a > 90 {
		return "A"
	} else if a <= 90 && a >= 60 {
		return "B"
	} else {
		return "Fail"
	}
}
func add(a, b int) int {
	return a + b
	//fmt.Println(a + b)
}

func swap(a, b string) (string, string) {
	return b, a
}
