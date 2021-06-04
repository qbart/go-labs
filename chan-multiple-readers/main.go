package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	dur := time.Duration(1000)
	go func(c <-chan int) {
		for {
			time.Sleep(dur * time.Millisecond)
			x := <-c
			fmt.Println("A", x)
		}
	}(ch)

	go func(c <-chan int) {
		time.Sleep(dur * time.Millisecond)
		x := <-c
		fmt.Println("B", x)
	}(ch)

	go func(c <-chan int) {
		time.Sleep(dur * time.Millisecond)
		x := <-c
		fmt.Println("C", x)
	}(ch)

	time.Sleep(50 * time.Millisecond)
	//for i := 0; i < 1000; i++ {
	ch <- 1
	fmt.Println("1.")
	ch <- 2
	fmt.Println("2.")
	ch <- 3
	fmt.Println("3.")
	ch <- 4
	fmt.Println("4.")
	ch <- 5
	fmt.Println("5.")
	ch <- 6
	fmt.Println("6.")

	fmt.Println("for {}")
	for {
	}
}
