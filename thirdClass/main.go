package main

//defer,panic,recover,struct,slice
import (
	"fmt"
	"sort"
)

// func mayPanic() {
// 	defer fmt.Println("defer in mayPanic")
// 	panic("oops!")
// }

// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered in main:", r)
// 		}
// 	}()

// 	mayPanic()

// 	fmt.Println("After mayPanic") // This will not run
// }

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered inside mayPanic:", r)
		}
	}()

	panic("oops!")
}

func main() {
	// mayPanic()
	// //panic("oops!")
	// fmt.Println("After mayPanic")

	// x := 5
	// defer fmt.Println(x)
	// x = 10

	structSample()

	//array
	// var a [2]string
	// a[0] = "nikunj"
	// fmt.Println(a)

	// //slice
	// var num = [6]int{0, 1, 2, 3, 4, 56}
	// var s []int = num[1:3]
	// fmt.Println(s)

	// var sl [5]int
	// sl[0] = 12
	// fmt.Println(sl)

	// //m:=make([]int,5)
	// //var s[]int
	// m := []int{}
	// m = append(m, 1)
	// fmt.Println(m)
	//fmt.Printf("%t", m)

	new := make([]int, 4)

	new[0] = 234
	new[1] = 945
	new[2] = 465
	new[3] = 867
	//new[4] = 777

	new = append(new, 555, 666, 321)
	sort.Ints(new)
	fmt.Println(new)

	//fmt.Println(sort.Ints(new))

}

func structSample() {
	type rectangle struct {
		length int
		breath int
	}
	var dimension = rectangle{1, 2}
	// fmt.Println(dimension)
	// fmt.Println(dimension.breath, dimension.length)
	// dimension.length = 89
	// fmt.Println(dimension.breath, dimension.length)
	fmt.Printf("%+v", dimension) //// i want my output as {length=89,breath=2}
}
