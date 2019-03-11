package main

import (
	"context"
	"fmt"
	"log"
	"net"

	weather "github.com/wwyiwzhang/gopumpkin/src/lessons/protobuf-grpc/weather"
	"google.golang.org/grpc"
)

type weatherServer struct {
}

func (w weatherServer) Inquire(ctx context.Context, wi *weather.WeatherInquiry) (*weather.WeatherReport, error) {
	var wt string

	switch wi.Date {
	case "today":
		wt = "cloudy"
	case "yesterday":
		wt = "rainy"
	case "tomorrow":
		wt = "rainbow"
	default:
		return nil, fmt.Errorf("I have no idea!")
	}
	return &weather.WeatherReport{
		Date:    wi.Date,
		Weather: wt,
	}, nil
}

func main() {

	srv := grpc.NewServer()
	var weathers weatherServer
	weather.RegisterWeathersServer(srv, weathers)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not listen on :8080, %v", err)
	}
	log.Fatal(srv.Serve(l))
}
