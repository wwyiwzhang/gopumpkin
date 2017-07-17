package main

import "fmt"

func main() {
  // if elif else
  x := 5
  if x == 1 {
    fmt.Println("x is 1")
  } else if x > 1 && x < 5 {
    fmt.Println("x is between 2 and 4")
  } else {
    fmt.Println("x is greater or equal to 5")
  }

  // for loop
  y := 10
  for y > 0 {
    fmt.Println("The number of y is ", y)
    y--
  }

  for z := 0; z <= 5; z++ {
    fmt.Println("The value of z is ", z)
  }

  // for loop is the while loop in C or Python

  // switch
  switch {
  case x == 1:
    fmt.Println("x is 1")
  case x > 1 && x <5:
    fmt.Println("x is between 2 and 4")
  case x >= 5:
    fmt.Println("x is greater or equal to 5")
  }

  // array
  var newArray[3] float64
  newArray[0] = 1.2
  newArray[1] = 1.3
  newArray[2] = 1.4
  fmt.Println("The first element of the array is", newArray[0])
  fmt.Println("The second element of the array is", newArray[1])
  fmt.Println("The third element of the array is", newArray[2])
  // slice the array
  fmt.Println("From the second to the last element of the array is", newArray[1:])

  // slice
  // []T where T is the type
  newArray1 := []float64 {1.2, 1.3, 1.4}
  // append to an existing slice
  // append(newArray1, 1.5) will give warning "evaludated but not used"
  newArray1 = append(newArray1, 1.5)
  fmt.Println("After appending 1.5 to the array, the array has become", newArray1)
  fmt.Println("The capacity of array is", cap(newArray1))

  // make can be used to create slice
  newArray2 := make([]float64, 0)
  newArray2 = append(newArray2, 1.2)
  newArray2 = append(newArray2, 1.3)
  newArray2 = append(newArray2, 1.4)
  fmt.Println("The array created using make is", newArray2)
  fmt.Println("The capacity of array is", cap(newArray2))

  newArray3 := make([]float64, 0, 4)
  newArray3 = append(newArray3, 1.2)
  newArray3 = append(newArray3, 1.3)
  // will throw an error if keep appending element to this slice
  newArray3 = append(newArray3, 1.4)
  fmt.Println("The array created using make is", newArray3)
  fmt.Println("The capacity of array is", cap(newArray3))

  // similar to Python enumerate function
  for ind, val := range newArray3 {
    fmt.Println("The ind and value is", ind, val)
  }

  // map key value
  // define a key value map with key as string and value as integer
  newMap := make(map[string] int)
  // insert some key value pairs
  newMap["first"] = 1
  newMap["second"] = 2
  newMap["third"] = 3
  fmt.Println("The key value map is", newMap)
  // remove some key value pairs
  delete(newMap, "third")
  fmt.Println("The key value map is", newMap)
  // add some key value pairs
  newMap["third"] = 3
  fmt.Println("The key value map is", newMap)
  // change the value for some key
  newMap["third"] = 4
  fmt.Println("The key value map is", newMap)
}
