package main

import "fmt"

func main() {
	tc1 := TempConvertor{18, "cel"}
	fmt.Println("fahrenheit is", tc1.convertor())
	tc2 := TempConvertor{100, "fah"}
	fmt.Println("celcius is", tc2.convertor())
	tc3 := TempConvertor{100, "fap"}
	fmt.Println("celcius is", tc3.convertor())
}

type TempConvertor struct {
	x float64
	y string
}

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
