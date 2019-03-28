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
	offSet := flag.Int("o", 0, "offset value")
	limit := flag.Int("m", 20, "query limit with max value of 50")
	fileName := flag.String("f", "", "name of the file to save records")
	flag.Parse()

	resp, err := yelper.Client().Get(
		yelpBusinessURL, map[string]string{
			"term":     *businessType,
			"location": *location,
			"limit":    fmt.Sprintf("%d", *limit),
			"offset":   fmt.Sprintf("%d", *offSet),
		})
	if err != nil {
		log.Fatalf("program failed to get from Yelp developer API, exiting...\n")
	}
	b, err := resp.Decode(yelper.Businesses{})
	if err != nil {
		log.Fatalf("program failed to decode businesses, %v\n", err)
	}
	// save business to CSV
	cb := b.ConvertItem()
	f, err := yelper.CreateORExists(*fileName)
	if err != nil {
		log.Fatalf("could not open file: %s to save data, %v\n", *fileName, err)
	}
	w := yelper.NewCSVWriter(f)
	w.WriteAll(cb)
	if err != nil {
		log.Fatalf("program faild to save businesses to CSV, %v\n", err)
	}
}
