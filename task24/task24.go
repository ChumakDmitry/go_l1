package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func (p *point) GetX() float64 {
	return p.x
}

func (p *point) GetY() float64 {
	return p.y
}

func Init(x, y float64) *point {
	p := point{
		x: x,
		y: y,
	}
	return &p
}

func main() {
	a := Init(1, 1)
	b := Init(2, 2)

	res := math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))

	fmt.Printf("Distance = %v", res)
}
