// package main

//concurrency
// import (
// 	"fmt"
// 	"time"
// )

// var count = 0

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			count++
// 		}()
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(count) // NOT 1000,we will use mutex.lock to solve it?
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(2) // we have 2 goroutines

	// go func() {
	// 	fmt.Println("Task 1 running")
	// 	wg.Done() // goroutine completed
	// }()

	// go func() {
	// 	fmt.Println("Task 2 running")
	// 	wg.Done() // goroutine completed
	// }()

	// wg.Wait() // wait for both goroutines
	// fmt.Println("All tasks finished")

	worker(1)
	go worker(2)
	time.Sleep(3 * time.Second)

	// 	var mu sync.Mutex
	// go func() {
	//     mu.Lock()   // can use directly because all goroutines share same variable
	//     mu.Unlock()
	// }()

}

func worker(id int) {
	for i := 0; i < 5; i++ {
		fmt.Println("worker,", id, "i :", i)
		time.Sleep(1 * time.Second)
	}

}
