package main

import (
	"fmt"
)

func main() {

	// res1 := sumEven(5)
	// fmt.Println("Sum of Even Numbers Between 1 and N", res1)

	// res2 := reverse(1234)
	// fmt.Println("Reverse a Number", res2)

	// res3 := countDigit(1234)
	// fmt.Println("Count the Digits in a Number", res3)

	// res4 := fibonacci(6)
	// fmt.Println("Print Fibonacci Series up to N Terms", res4)

	// res5 := prime(2)
	// fmt.Println("Check if a Number is Prime", res5)

	// pyramid(6)

	// res7 := sumDig(1234)
	// fmt.Println("Find the Sum of Digits", res7)

	// res8 := facrorial(5)
	// fmt.Println("Factorial Using Loop", res8)

	// res9 := reverseStr("nikunj")
	// fmt.Println("Reverse a String (without using library functions)", res9)

	// diamond(6)
	//fibo(6)
	// var n int = 0123
	// fmt.Println(n, "vghhj")

	//fmt.Println(grade(78))
	//difer_check()
	//var x int = 5
	pointer()

}
func pointer() {
	var x int = 5
	// fmt.Println(x)
	// fmt.Println(&x)
	p := &x
	fmt.Printf("type=%T", "\n valve=%v", p, p)
	fmt.Println(*p)
	fmt.Printf("check %x=", x)
}

func grade(marks int) string {
	grade := ""
	switch marks {
	// case marks > 90:
	// 	grade = "A"
	// case marks < 90 && marks > 60:
	// 	grade = "B"

	case 78:
		grade = "A===="
	default:
		grade = "F"

	}
	return grade
}

func difer_check() {
	fmt.Println("After1")
	fmt.Println("Befor")

	defer fmt.Println("After2")
	fmt.Println("Befor")
}

// Sum of Even Numbers Between 1 and N
func sumEven(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

// Reverse a Number
func reverse(n int) int {
	newN := 0
	for n > 0 {
		digit := n % 10
		newN = newN*10 + digit
		n = n / 10
	}
	return newN
}

// Count the Digits in a Number
func countDigit(n int) int {
	count := 0
	for n > 0 {
		count += 1
		n = n / 10
	}
	return count
}

// Print Fibonacci Series up to N Terms
func fibonacci(n int) []int {
	res := []int{0, 1}
	if n == 0 {
		return nil
	} else if n == 1 {
		return []int{0}
	}

	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])

		//res[i] = res[i-1] + res[i-2]
	}
	return res
}

// Check if a Number is Prime
func prime(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Print a Pyramid Pattern
// diamond, chri tree
func pyramid(n int) {
	fmt.Println("printing pyramid::")
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func diamond(n int) {
	fmt.Println("printing diamond::")
	for i := 1; i <= n; i++ {
		for k := n - i; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	for i := n - 1; i > 0; i-- {
		for k := 0; k < n-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

// Find the Sum of Digits
func sumDig(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n = n / 10
	}
	return sum
}

// Factorial Using Loop
func facrorial(n int) int {
	res := 1 //0 haga then 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// Reverse a String (without using library functions)
func reverseStr(str string) string {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	return newStr
}

// fib 2nd way
func fibo(n int) {
	prev1 := 0
	prev2 := 1
	fmt.Println(prev1, prev2)
	for i := 2; i < n; i++ {
		fmt.Print(prev1 + prev2)
		temp := prev2
		prev2 = prev1 + prev2
		prev1 = temp
	}
}
