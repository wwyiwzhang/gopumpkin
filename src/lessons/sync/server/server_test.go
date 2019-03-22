package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	b, err := ioutil.ReadFile("example.json")
	if err != nil {
		t.Fatalf("could not read file, %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("could not create test request, %v", err)
	}
	rec := httptest.NewRecorder()
	handlerFunc(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("incorrect status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response payload, %v", err)
	}
	if exp := "received counter: 1"; string(data) != exp {
		t.Fatalf("expected :%s, got: %s", exp, string(data))
	}
}

func BenchmarkHandler(b *testing.B) {
	b.StopTimer()

	d, err := ioutil.ReadFile("example.json")
	if err != nil {
		b.Fatalf("could not read file, %v", err)
	}
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", bytes.NewReader(d))
		if err != nil {
			b.Fatalf("could not create test request, %v", err)
		}
		rec := httptest.NewRecorder()
		b.StartTimer()
		handlerFunc(rec, req)
		b.StopTimer()
	}
	b.ReportAllocs()
}
