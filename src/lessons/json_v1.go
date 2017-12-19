package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	// "reflect"
)

// get some data from recipe puppy
type Recipe struct {
	Title string
	Href string
	Ingredients string
	Thumbnail string
}

type Jsondata struct {
	Title string 
	Version float64
	Href string
	Results []Recipe
}

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func ParseJson(data string, mapper interface{}) ([]interface {}) {
	err := json.Unmarshal([]byte(data), &mapper)
	if(err != nil) {
		panic(err)
	}
	result := mapper.(map[string]interface{})["results"]
	fmt.Println(result)
	return result.([]interface{})
}

func main() {
	var url = "http://www.recipepuppy.com/api/?i=onions,garlic&q=omelet"
	var data, err = Request(url) // the return type is a string
	if err != nil {
		fmt.Println(err)
	}
	mapper := map[string]interface{}{}
	parsed := ParseJson(data, mapper)
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