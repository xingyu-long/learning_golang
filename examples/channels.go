package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// firstExample()
	// secondExample()
	// thirdExample()
	// fourthExample()
	fifthExample()
}

func firstExample() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)

	go func() {
		// pass this as copy.
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()
}

func secondExample() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(2)

		go func() {
			// pass this as copy.
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			// unbuffered channel: Only one message can stay at this channel.
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

func thirdExample() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	// receive from channel only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send to channel only
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}

func fourthExample() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// NOTE: for this `range` loop, just single value we got
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		// it's same with above code.
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // empty allocation.

func fifthExample() {
	// Use select switch to terminate the channel.
	// also use empty struct
	go logger()

	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is starting"}
	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time, entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
