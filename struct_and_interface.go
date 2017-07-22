package main

import (
	"fmt"
	"math"
)

func main() {
	tc1 := TempConvertor{18, "cel"}
	fmt.Println("fahrenheit is", tc1.convertor())
	tc2 := TempConvertor{100, "fah"}
	fmt.Println("celcius is", tc2.convertor())
	tc3 := TempConvertor{100, "fap"}
	fmt.Println("celcius is", tc3.convertor())
	cir := Circle{1.5}
	squ := Square{2}
	fmt.Println("The area of a circle is", cir.getarea())
	fmt.Println("The area of a square is", squ.getarea())
	fmt.Println("The area of the object is", printArea(cir))
	fmt.Println("The area of the object is", printArea(squ))
}

type TempConvertor struct {
	x float64
	y string
}

// (tc *TempConvertor) is a function receiver
// panic not recommended to use
func (tc *TempConvertor) convertor() float64 {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("the only two options are 'fah' and 'cel'")
		}
	}()

	if tc.y == "cel" {
		return tc.x*1.8 + 32
	} else if tc.y == "fah" {
		return (tc.x - 32) / 1.8
	} else {
		panic("unknown type!")
	}
}

// interface
// a very classic example
type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (c Circle) getarea() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) getarea() float64 {
	return s.side * s.side
}

type Geometry interface {
	getarea() float64
}

func printArea(g Geometry) float64 {
	return g.getarea()
}
