package main

import (
	"net"
	"flag"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "8800", "port")
var firstname = flag.String("firstname", "", "first name")
var lastname = flag.String("lastname", "", "last name")
var order = flag.String("order", "", "order")

type Response struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Order string `json:"order"`
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Failed to connect: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	go sendData(conn, done)
	// go getResponse(conn, done)
	fmt.Println(<-done)
	// fmt.Println(<-done)
}

func getResponse(conn net.Conn, done chan string) {
	statuscode, err_r := ioutil.ReadAll(conn)
	if err_r != nil {
		fmt.Println("error reading: ", err_r)
		panic(err_r)
	}
	i, _ := strconv.Atoi(string(statuscode))
	fmt.Println("Status code from server is: ", i)
	done <- "Read"
}

func sendData(conn net.Conn, done chan string) {
	r := Response{
		Firstname: *firstname,
		Lastname: *lastname,
		Order: *order,
	}
	b, errs := json.Marshal(r)
	if errs != nil {
		fmt.Println("Failed to serialize data")
		return
	}
	_, err := conn.Write(b)
	if err != nil {
		fmt.Println("Failed to make a request to server with data: ", string(b))
		return
	}
	fmt.Println("Succeeded to make a request to server with data: ", string(b))
	done <- "Sent"
}

