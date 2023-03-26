package main

import (
	"fmt"
	"time"
)

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func concurrency(n int, result chan int) {
	result <- fibo(n)
}

func main() {
	t1 := time.Now().UnixMilli()
	result := fibo(40)
	fmt.Println(result)
	// result1 := make(chan int)
	// result2 := make(chan int)
	// go concurrency(38, result1)
	// go concurrency(39, result2)
	// fmt.Println(<-result1 + <-result2)
	t2 := time.Now().UnixMilli()
	fmt.Printf("Time: %v", t2-t1)
}
