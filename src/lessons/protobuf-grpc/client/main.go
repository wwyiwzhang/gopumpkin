package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	weather "github.com/wwyiwzhang/gopumpkin/src/lessons/protobuf-grpc/weather"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "not enough argument!")
		os.Exit(1)
	}
	if flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, "too many arguments: expected 1, found %d", flag.NArg())
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to weather service: %v", err)
	}
	client := weather.NewWeathersClient(conn)

	cmd := flag.Arg(0)

	weatherInquiry := weather.WeatherInquiry{
		Date: cmd,
	}

	wr, err := client.Inquire(context.Background(), &weatherInquiry)
	if err != nil {
		log.Fatalf("weather inquire failed, %v", err)
	}
	log.Printf("Weather report for %s: %s", weatherInquiry.Date, wr.Weather)
}
