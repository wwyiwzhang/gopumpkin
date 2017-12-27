package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// get some data from recipe puppy
type Recipe struct {
	Title string `json:"title"`
	Href string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail string `json:"thumbnail"`
}

type Jsondata struct {
	Title string `json:"title"`
	Version float64 `json:"version"`
	Href string `json:"href"`
	Results []Recipe `json:"results"`
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

func ParseJson(data string, mapper Jsondata) ([]Recipe) {
	err := json.Unmarshal([]byte(data), &mapper)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(mapper)
	result := mapper.Results
	fmt.Println(result)
	return result
}

func main() {
	var url = "http://www.recipepuppy.com/api/?i=onions,garlic&q=omelet"
	var data, err = Request(url) // the return type is a string
	if err != nil {
		fmt.Println(err)
	}
	var mapper Jsondata
	parsed := ParseJson(data, mapper)
	// fmt.Println(parsed)
	for i := range parsed {
		// needs to do a type assertion here
		recipe_mapper := parsed[i]
		fmt.Println(recipe_mapper.Title)
		fmt.Println(recipe_mapper.Ingredients)
		fmt.Println(recipe_mapper.Href)
	}
}