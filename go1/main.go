package main

import (
	"fmt"
)

// const (
// 	q = iota << 1
// 	w
// 	e
// )

func anything() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
}

func sum(n ...int) int {
	s := 0
	for _, x := range n {
		s += x
	}
	return s
}

func handle() (string, bool) {
	return "", false
}

func main() {
	a := make([]int, 5, 10)
	for i := 0; i < 5; i++ {
		a[i] = i + 1
	}

	// 0  1  2  3  4
	// a = []int{1, 2, 3, 4, 5}
	// a = a[:i+copy(a[i:], a[i+1:])]
	// a = append(a[:i], a[i+1:]...)
	// a = append(a[:2], a[3:]...)

	fmt.Println(a)
	a = make([]int, 0, 7)

	for i := 0; i < 100; i++ {
		a = append(a, 2)
		fmt.Println(i, "cap=", cap(a))
	}

	fmt.Println(cap(a))
	fmt.Println(len(a))

	b := make(map[string]string)
	b["key"] = "value"

	x, ok := b["key"]
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("none")
	}
	// delete(a, "key")

	for k, v := range b {
		fmt.Println(k, v)
	}

	for _, elem := range b {
		fmt.Println(elem)
	}

	d := 1
	p := &d

	*p = 2
	// d := &alg.Point{X: 10, Y: 2}
	// d.X = 20
	fmt.Println(d)

	// anything()
	// f, err := file.Open("/tmp/random")
	// if err != nil {
	// 	panic("Fail to open")
	// }
	// defer f.Close()

	//processing

	// d := alg.Point{X: 10, Y: 2}
	// x, y := d.Swap()
	// fmt.Println(x, y)
	// fmt.Println(q, w, e)

	// a := []int{1, 2, 3}
	// for x == 2 {
	// }
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// for _, elem := range a {
	// 	fmt.Println(elem)
	// }
}
