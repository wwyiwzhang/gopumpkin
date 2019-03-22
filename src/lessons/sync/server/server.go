package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Counter struct {
	ID int `json:"id"`
}

var s = sync.Pool{
	New: func() interface{} {
		return new(Counter)
	},
}

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	c := s.Get().(*Counter)
// 	defer s.Put(c)
// 	c.ID = 0
// 	err := json.NewDecoder(r.Body).Decode(c)
// 	if err != nil {
// 		log.Fatalf("could not decode object %v", err)
// 	}
// 	fmt.Fprintf(w, "received counter: %d", c.ID)
// }

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	c := &Counter{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		log.Fatalf("could not decode object %v", err)
	}
	fmt.Fprintf(w, "received counter: %d", c.ID)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting server at :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
