package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func main() {
	// set environment variables
	os.Setenv("Foo", "Bar")
	os.Setenv("Pumpkin", "Wow")
	fmt.Println("Foo: ", os.Getenv("Foo"))
	fmt.Println("Pumpkin: ", os.Getenv("Pumpkin"))
	// what if the env doesn't exist, it should return an empty string
	fmt.Println("Home :", os.Getenv("Home"))
	home := os.Getenv("Home")
	fmt.Println("The type of environment variable Home is: ", reflect.TypeOf(home))
	if home == "" {
		fmt.Println("It is an empty string")
	}
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		fmt.Printf("Key: %s, Value: %s", kv[0], kv[1])
	}
	fmt.Println("\n")
	// regex
	regexstring := "hello World bigworld"

	// match string 
	r, _ := regexp.Compile("e[a-z]+o")
	fmt.Println(r.MatchString(regexstring))
	fmt.Println(r.FindString(regexstring))
	
	// match string case insensitive
	ri, _ := regexp.Compile("(?i)w[a-z]+d")
	fmt.Println(ri.MatchString(regexstring))
	fmt.Println(ri.FindString(regexstring))	
	fmt.Println(ri.FindStringIndex(regexstring))	
	// get all matches

	ra, _ := regexp.Compile("(?i)w[a-z]+d")
	fmt.Println(ra.FindAllString(regexstring, 2)) // return as []string
	fmt.Println(reflect.TypeOf(ra.FindAllString(regexstring, 2)))
}