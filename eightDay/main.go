package main

//channel,concurrency
import (
	"fmt"
	"time"
)

// var mu sync.Mutex
// var wg sync.WaitGroup
// var balance = 0

// func deposit(amount int) {
// 	//mu.Lock()
// 	balance += amount
// 	//mu.Unlock()
// 	wg.Done()
// 	//fmt.Println(balance)//go run --race .
// }

func main() {
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go deposit(100)
	// }

	// wg.Wait()
	// fmt.Println("Final Balance:", balance)

	// ch := make(chan int)
	// ch <- 10
	// fmt.Println(<-ch)//unbuffered,deadlock

	// ch := make(chan int)
	// defer func() {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// }()
	// go func() {
	// 	//ch:=make(chan int)
	// 	ch <- 10
	// 	ch <- 9

	// }()
	// defer func() {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// }()
	//time.Sleep(time.Second)

	// ch := make(chan int, 3) // 3 is the limit ,buffered
	// //fmt.Println(<-ch)
	// func() {
	// 	//ch := make(chan int, 3)
	// 	//ch:=make(chan int)
	// 	ch <- 10
	// 	ch <- 9
	// 	ch <- 10
	// 	//ch <- 9

	// }()
	// func() {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// 	//fmt.Println(<-ch)
	// }()

	// ch := make(chan int)
	// go workerW(ch)

	// workerR(ch) //normal read
	ch := make(chan int)
	done := make(chan bool)

	go writer(ch)
	go reader(ch, done)

	<-done // wait for reader to finish
	fmt.Println("Main finished.")
}
func workerW(ch1 chan int) {
	//ch1<-10
	for i := 0; i < 5; i++ {
		// if i == 2 {
		// 	close(ch1)
		// }
		ch1 <- i

	}
	close(ch1)
}

func workerR(ch1 chan int) {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-ch1)
	// }

	// for value := range ch1 {
	// 	fmt.Println(value)
	// }

	select {
	case val := <-ch1:
		fmt.Println("Received:", val)
		//case :

	}

}

func writer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond) // simulate work
	}
	close(ch) // important: signal reader to stop
}

func reader(ch chan int, done chan bool) {
	for {
		select {

		// case 1: receive value
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, reader exiting.")
				done <- true
				return
			}
			fmt.Println("Received:", val)

		// case 2: timeout if no value arrives in time
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout: no data received.")
			done <- true
			return
		}
	}
	// for {
	// 	select {
	// 	case val, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("channel closed,ended")
	// 			return
	// 		}
	// 		fmt.Println(val)
	// 	}
	// }
}

// normal func ka go
//close(ch)
//range and select in channel
//defer mei go?
