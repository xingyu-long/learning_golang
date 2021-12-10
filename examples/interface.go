package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	// why we need the `&`?
	// var ic Incrementer = &myInt
	ic := &myInt
	fmt.Printf("%T", ic)
	for i := 0; i < 10; i++ {
		fmt.Println(ic.Increment())
	}
	// NOTE: Why it didn't work? 
	// because I wrote `closer()` for Closer.
	var wc WriterCloser = myWriterCloser{}
	// if one of receivers needs *, so we should pass &.
	// Method set of value is all methods with value receivers.
	// Method set of pointer is all methods, regardless of receiver type.
	// wc := myWriterCloser{}
	wc.Write([]byte("Hello"))
	wc.Close()
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// another example of interface: Incrementer.
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// WriterCloser example
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mvc myWriterCloser) Write(data []byte) (int, error) {
	fmt.Println("Here is a write part")
	return 0, nil
}

func (mvc myWriterCloser) Close() error {
	fmt.Println("Here is a close part")
	return nil
}
