/*
March 21, 2019, writing small examples of functional options
*/

package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	name   string
	client *http.Client
	token  string
}

type Option func(*Client)

func NewClient(n string, options ...Option) *Client {
	c := &Client{
		name:   n,
		client: &http.Client{},
		token:  "",
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func newLogger(t string) Option {
	return func(c *Client) {
		c.token = t
	}
}

func main() {
	// default
	d := NewClient("default")
	fmt.Printf("default client: %v\n", d)

	// functional options
	s := NewClient("special", newLogger("specialtoken"))
	fmt.Printf("default client: %v\n", s)
}
