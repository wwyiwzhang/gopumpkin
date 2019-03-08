package main

import (
	"fmt"
	"strings"
)

type Enricher struct {
}

func (e *Enricher) enrich(str string) string {
	return "===" + str + "==="
}

type Standardizer struct {
	String string
}

func (s *Standardizer) standardize(str string) string {
	return strings.ToUpper(str)
}

type Datapipeline struct {
	Enricher
	Standardizer
}

func main() {
	d := new(Datapipeline)
	greeting := "hello world"
	std_greeting := d.enrich(greeting)
	fmt.Println(std_greeting)
	enriched_std_greeting := d.standardize(std_greeting)
	fmt.Println(enriched_std_greeting)
}
