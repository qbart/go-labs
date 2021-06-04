package main

// test all combinations:
// - buffered ch without closing channel c
// - unbuffered ch without closing channel c
// - buffered ch and closing channel c
// - unbuffered ch and closing channel c

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // buffered
	// ch := make(chan int) // unbuffered

	// reader - allow only reads from channel
	go func(c <-chan int) {
		for {
			time.Sleep(2000 * time.Millisecond)
			x := <-c
			fmt.Println("R", x)
		}
	}(ch)

	// writer - allow only writes to channel
	go func(c chan<- int) {
		x := 1
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("W", x)
			c <- x
			x++
			if x > 5 {
				// close(c) // uncomment to see what happens
				break
			}
		}
	}(ch)

	// prevent exit
	for {
	}
}
