package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func(c <-chan int) {
		for {
			time.Sleep(200 * time.Millisecond)
			// detecting if channel closed, rather than relying on default value
			x, ok := <-c
			if !ok {
				fmt.Println("Channel closed, sending quit")
				quit <- true
				break
			}
			fmt.Println("R", x, ok)
		}
	}(ch)

	go func(c chan<- int) {
		x := 1
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("W", x)
			c <- x
			x++
			if x > 5 {
				close(c)
				break
			}
		}
	}(ch)

	// prevent exit
MAIN_LOOP:
	for {
		select {
		case <-quit:
			break MAIN_LOOP
		}
	}

	fmt.Println("Bye")
}
