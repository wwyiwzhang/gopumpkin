package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func sendRequest(ctx context.Context, id int) {
	select {
	case <-ctx.Done():
		return
	default:
		var jsonStr = []byte(fmt.Sprintf(`{"id": %d}`, id))
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", bytes.NewBuffer(jsonStr))
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			log.Fatalf("could not post data: %s, %v", string(jsonStr), err)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	i := 0
	for {
		go sendRequest(ctx, i)
		time.Sleep(100 * time.Millisecond)
		i++
	}
}
