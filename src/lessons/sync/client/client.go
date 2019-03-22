package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sendRequest(ctx context.Context, id int) {
	select {
	case <-ctx.Done():
		return
	default:
		var jsonStr = []byte(fmt.Sprintf(`{"id": %d}`, id))
		req, err := http.NewRequest("POST", "http://localhost:8000", bytes.NewBuffer(jsonStr))
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			log.Fatalf("could not post data: %s, %v", string(jsonStr), err)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	i := 0
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL)
	for {
		select {
		case <-sig:
			fmt.Println("program killed")
			cancel()
			os.Exit(0)
		default:
			go sendRequest(ctx, i)
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}
