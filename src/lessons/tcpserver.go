package main

import (
	"net"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "8800", "port")

type Order struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Order string `json:"order"`
}

type Response struct {
	Statuscode int `json:"statuscode"`
	Message string `json:"message"`
}

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting: ", err)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	data, err := ioutil.ReadAll(conn) // return as bytes array
	var statuscode int
	var message string
	if err != nil {
		fmt.Println("reading err: ", err)
		statuscode = 400
		message = "Failed"
	} else {
		var order Order
		json.Unmarshal(data, &order)
		fmt.Println("Received: Firstname: " + order.Firstname, ", " + "Lastname: ", order.Lastname, ", " + "Order: ", order.Order)
		statuscode = 200
		message = "Success"
	}
	response := Response{
		Statuscode: statuscode,
		Message: message,
	}
	m, err_m := json.Marshal(response)
	if err_m != nil {
		fmt.Println("Failed to serialize response")
	}
	_, err_r := conn.Write(m)
	if err_r != nil {
		fmt.Println("Failed to send status code to client")
	}
}