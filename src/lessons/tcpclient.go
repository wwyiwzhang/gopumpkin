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
	conn, err := net.Dial("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Failed to connect: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	sendDataNgetResponse(conn)
	conn.Close()
}

func sendDataNgetResponse(conn net.Conn) {
	order := Order{
		Firstname: *firstname,
		Lastname: *lastname,
		Order: *order,
	}
	b, errs := json.Marshal(order)
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

	conn.(*net.TCPConn).CloseWrite()

	message, err_r := ioutil.ReadAll(conn)

	var r Response

	json.Unmarshal(message, &r)

	if err_r != nil {
		fmt.Println("error reading: ", err_r)	
	} else {
		if r.Statuscode == 200 {
			fmt.Println("Server succeeded to respond to the request with status code: ", strconv.Itoa(r.Statuscode))
		} else {
			fmt.Println("Server failed to respond to the request with status code: ", strconv.Itoa(r.Statuscode))
		}
	}
}

