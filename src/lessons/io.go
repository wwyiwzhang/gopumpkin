package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {

	// TeeReader example
	// Takes a reader and writer and return a reader
	// it will output to writer for whatever it reads in
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not stat standard input: %v", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "reading data from terminal, not implemented!")
		os.Exit(1)
	} else {
		fmt.Println("reading data from pipe")
	}
	s := bufio.NewScanner(io.TeeReader(os.Stdin, os.Stdout))
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer close(c)
		defer wg.Done()
		for s.Scan() {
			c <- s.Text()
		}
		if err := s.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to read from standard input: %v", err)
		}
	}()
	wg.Add(2)
	go func() {
		defer wg.Done()
		for c != nil {
			_, ok := <-c
			if !ok {
				c = nil
			}
		}
	}()
	wg.Wait()

	// MultiWriter example
	// Writes to multiple writers: Stdout and a text file
	d := strings.NewReader("hello world, learning io library is fun\n")
	f, err := os.Create("output.txt")
	// defer f.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file for writing: %v", err)
		os.Exit(1)
	}
	writers := io.MultiWriter(os.Stdout, f)
	_, err = io.Copy(writers, d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not copy from reader to writers: %v", err)
	}
}
