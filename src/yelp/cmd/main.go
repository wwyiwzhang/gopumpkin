package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
	yelper "github.com/wwyiwzhang/gopumpkin/src/yelp/api"
)

const (
	yelpBusinessURL = "https://api.yelp.com/v3/businesses/search"
)

func main() {

	businessType := flag.String("t", "restaurants", "business type")
	location := flag.String("l", "sf", "business location")
	offSet := flag.Int("o", 20, "offset value")
	limit := flag.Int("m", 20, "query limit")
	flag.Parse()

	client := yelper.NewYelpHTTPClient()
	log.Infoln("created client to connect to Yelp developer API")

	resp, err := client.Get(yelpBusinessURL, map[string]string{
		"term":     *businessType,
		"location": *location,
		"limit":    fmt.Sprintf("%d", *offSet),
		"offset":   fmt.Sprintf("%d", *limit),
	})
	if err != nil {
		log.Fatalf("program failed to get from Yelp developer API, exiting...\n")
	}
	var business yelper.Businesses
	b, err := resp.Decode(business)
	if err != nil {
		log.Fatalf("program failed to decode businesses, %v\n", err)
	}
	// save business to CSV
	cb := b.ConvertItem()
	f, err := yelper.CreateORExists("yelp.csv")
	if err != nil {
		log.Fatalf("could not open file to save data, %v\n", err)
	}
	w := yelper.NewCSVWriter(f)
	err = w.Write(cb)
	if err != nil {
		log.Fatalf("program faild to save businesses to CSV, %v\n", err)
	}

}
