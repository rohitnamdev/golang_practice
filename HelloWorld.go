package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// type Details interface {
// 	data()
// }

// func data() {
// 	fmt.Println("i m data")
// }

// type non struct {
// 	a string
// 	b int64
// 	c float64
// }

// type adhar struct{}

// func (*pan) atrangi() {
// 	fmt.Println("data from Pan...")
// }

// func add(a, b int64) int64 {
// 	return a + b
// }

// func (*adhar) darta() {
// 	fmt.Println("Mein toh nahi darta...")
// }

// type pan struct{}

// func (*pan) data() {
// 	fmt.Println("data from Pan...")
// }

// func main() {
// 	var a Details = &adhar{}
// 	a.data()
// 	a = &pan{}
// 	a.data()
// 	var b pan
// 	b.data() //invoking
// 	data()   //direct function calling
// }

// Prints numbers from 1-3 along with the passed string

func foo(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ": ", i)
		wg.Done()
	}
}

func main() {
	// fmt.Println("Main started...")
	// // Starting two goroutines
	// wg.Add(1)
	// go foo("1st goroutine")
	// wg.Wait() // First go routine
	// go foo("2nd goroutine")
	// wg.Wait()
	// // Wait for goroutines to finish before main goroutine ends
	// time.Sleep(time.Second)
	// fmt.Println("Main finished")
	fmt.Println("HelloWorld")
}

// func myFunc(waitgroup *sync.WaitGroup) {
// 	fmt.Println("Inside my goroutine")
// 	waitgroup.Done()
// }
// func main() {
// 	fmt.Println("Hello World")
// 	var waitgroup sync.WaitGroup
// 	waitgroup.Add(1)
// 	go myFunc(&waitgroup)
// 	waitgroup.Wait()
// 	fmt.Println("Finished Execution")

// }
