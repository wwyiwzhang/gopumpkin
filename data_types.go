package main

import "fmt"

func main() {
  // integer
  var expInteger int = 40
  fmt.Println(expInteger)
  expInteger2 := 40
  fmt.Println(expInteger2)
  // assign a new value to an existing variable
  expInteger2 = 50
  fmt.Println(expInteger2)

  // float
  var expFloat float64 = 1.334343
  fmt.Println(expFloat)
  expFloat2 := 1.334343
  fmt.Println(expFloat2)

  // string
  var expString string = "Hello World"
  fmt.Println(expString)
  expString2 := "Hello World"
  fmt.Println(expString2)
  fmt.Println(len(expString2))

  // boolean
  var isRaining bool = false
  fmt.Println(isRaining)

  // declare a constant
  const Pi float64 = 3.141516
  // round to the third decimal places
  fmt.Printf("%.3f \n", Pi)

  // declare multiple variables
  var (
    first = "one"
    second = "second"
    third = "third"
  )
  fmt.Println(first);
  fmt.Println(second);
  fmt.Println(third);
}
