package main

import (
	"net"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "8800", "port")

type Response struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Order string `json:"order"`
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
	if err != nil {
		fmt.Println("reading err: ", err)
		statuscode = 400
	} else {
		var r Response
		json.Unmarshal(data, &r)
		fmt.Println("Received: Firstname: " + r.Firstname, ", " + "Lastname: ", r.Lastname, ", " + "Order: ", r.Order)
		statuscode = 200
	}
	_, err_r := conn.Write([]byte("Status code: " + strconv.Itoa(statuscode)))
	if err_r != nil {
		fmt.Println("Failed to send status code to client")
	}
	fmt.Println("Sending status code: ", statuscode)
}