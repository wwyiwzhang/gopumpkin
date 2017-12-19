package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"reflect"
)

// get some data from recipe puppy

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func ParseJson(data string, mapper map[string]interface{}) ([]interface {}) {
	json.Unmarshal([]byte(data), &mapper)
	result := mapper["results"]
	return result.([]interface {})
}

func main() {
	var url = "http://www.recipepuppy.com/api/?i=onions,garlic&q=omelet"
	var data, err = Request(url) // the return type is a string
	if err != nil {
		fmt.Println(err)
	}
	var mapper map[string]interface{}
	parsed := ParseJson(data, mapper)
	fmt.Println(reflect.TypeOf(parsed))
	for i := range parsed {
		// needs to do a type assertion here
		recipe_mapper, ok := parsed[i].(map[string]interface{})
		if ok {
			fmt.Println(recipe_mapper["title"])
			fmt.Println(recipe_mapper["ingredients"])
			fmt.Println(recipe_mapper["href"])
		} else {
			fmt.Println("Failed to type assert item num %v", i+1)
		}
	}
}