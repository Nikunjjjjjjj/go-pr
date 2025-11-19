package main

//maps,function value, func inside func,closure part1
import (
	"fmt"
	"strings"
)

func main() {
	// nums := []int{1, 2, 3, 4, 17, 20, 97}
	// fmt.Printf("%T", nums)
	// for _, n := range nums {
	// 	fmt.Println(n)
	// }

	// type rectangle struct {
	// 	length int
	// 	breath int
	// }
	// m := make(map[string]rectangle)
	// m["rec1"] = rectangle{1, 2}
	// var m = map[string]rectangle{
	// 	"rec1": {1, 2},
	// 	"rec2": {3, 4},
	// }
	// m["rec3"] = rectangle{5, 6}
	// // fmt.Println(m)
	// // fmt.Printf("%T", m)
	// for i, j := range m {
	// 	fmt.Println(i, "gjghjgjhgjh", j)
	// }

	// m := make(map[string]int)
	// m["a"] = 10
	// fmt.Println(m["a"])

	// m["a"] = 100
	// fmt.Println(m["a"])

	// delete(m, "a")
	// fmt.Println(m)
	// fmt.Println(m["a"])

	// v, ok := m["a"]
	// fmt.Println(v, ok)

	// res := split("Blockchain is a decentralized, distributed ledger technology that allows digital information to be recorded, stored, and verified in a secure and transparent manner without the need for a central authority. It consists of a chain of blocks, each containing a list of transactions or data, which are cryptographically linked to the previous block, making the entire chain tamper-resistant and immutable. Each participant in the network, often referred to as a node, maintains a copy of the blockchain, ensuring consensus and trust among parties. The technology is best known as the underlying structure for cryptocurrencies like Bitcoin and Ethereum, but its applications extend far beyond digital money, including supply chain management, voting systems, identity verification, and smart contracts. By enabling peer-to-peer transactions that are transparent, verifiable, and resistant to fraud, blockchain has the potential to transform industries by reducing intermediaries, increasing efficiency, and providing new levels of security and accountability. Despite its promise, challenges such as scalability, energy consumption, regulatory uncertainty, and integration with existing systems remain, but ongoing research and innovation continue to expand its practical use cases and adoption across the globe. ")
	//fmt.Println(res)
	// mapping(res)
	// res := split("anshul is a anshul ")
	// fmt.Println("the word array : ", res[0])

	// countVowel("my name is nikunj is ")
	// fmt.Println(countEachVowel("my name is nikunj is "))

	//function value
	// mul := multiply
	// fmt.Println(mul(1, 2))

	// fmt.Println(processnumber(3, cube))
	// fmt.Println(processnumber(2, square))
	//fmt.Println(processnumber(2, multiply))

	// s := addr()
	// fmt.Println(s(5))
	// fmt.Println(s(6))
	// fmt.Println(s) //memory address
	// fmt.Println(addr()(6))

	fibonacci(6)
}

// Constructor
// func newRectangle(l, b int) rectangle {
// 	return rectangle{
// 		length:  l,
// 		breadth: b,
// 		area:    l * b,
// 	}
// }

func split(str string) []string {
	arr := []string{}
	//str=string.tolowercase
	var word string
	for i := 0; i < len(str); i++ {

		if str[i] != ' ' {
			word += string(str[i])
		} else {
			arr = append(arr, word)
			word = ""
		}

	}
	return arr
}

func mapping(arr []string) {
	m := make(map[string]int)
	for _, v := range arr {
		m[v] = m[v] + 1
	}
	fmt.Println(m)
}

// making.include
func countVowel(str string) {
	m := make(map[string]int)
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range str {
		//fmt.Println(string(v))
		for _, vowel := range vowels {
			if string(v) == string(vowel) {
				m[string(vowel)] += 1
				break
			}
		}
	}
	//fmt.Println(str, "cgcghg")
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(string(str[i]))
	// }
	fmt.Println(m)

}

func countEachVowel(s string) map[string]int {
	// Convert string to lowercase to make counting case-insensitive
	s = strings.ToLower(string(s))
	vowels := "aeiou"
	vowelCount := make(map[string]int)

	// Initialize the map
	// for _, v := range vowels {
	// 	vowelCount[v] = 0
	// // }

	// Count vowels
	for _, char := range s {
		if strings.Contains(vowels, string(char)) {
			vowelCount[string(char)]++
		}
	}

	return vowelCount
}

//string.fields

func multiply(x, y int) int {
	return x * y
}

func processnumber(x int, fun func(int) int) int {
	return fun(x)
}

func cube(x int) int {
	return x * x * x
}
func square(x int) int {
	return x * x
}

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci with closure
func fi() func(x, y int) int {
	//s := 0
	return func(x, y int) int {
		s := x + y
		return s
	}
}

// listnode se try
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func fibonacci(x int) {
	sum := fi()
	current := 0
	next := 1
	if x == 0 {
		return
	}
	if x == 1 {
		fmt.Println(0)
		return
	}
	fmt.Println(current, "\n", next)
	for i := 2; i < x; i++ {
		res := sum(current, next)
		fmt.Println(res)
		temp := next
		next = res
		current = temp
		//sum=fi()
	}
}
