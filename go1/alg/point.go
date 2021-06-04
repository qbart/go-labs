package alg

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%dx%d", p.X, p.Y)
}

// func (p *Point) Swap() {
// 	a := p.X
// 	p.X = p.Y
// 	p.Y = a
// }

func (p Point) Swap() (int, int) {
	return p.Y, p.X
}
