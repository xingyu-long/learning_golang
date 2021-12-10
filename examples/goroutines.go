package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// example 1:
	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

	// example 2:
	// msg := "Hello Again!"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Goodbye"
	// wg.Wait() // make sure main function won't exit too early.

	// example 3:
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go sayHelloWithCounter()
	// 	go increment()
	// }
	// wg.Wait()

	// example 4: Mutex
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // reader lock
		go sayHelloWithMutex()
		m.Lock()
		go incrementWithMutex()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello Go!")
}

func sayHelloWithCounter() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func sayHelloWithMutex() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func incrementWithMutex() {
	counter++
	m.Unlock()
	wg.Done()
}
