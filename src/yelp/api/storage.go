package yelper

import (
	"bufio"
	"encoding/csv"
	"os"
)

func NewCSVWriter(f *os.File) *csv.Writer {
	return csv.NewWriter(bufio.NewWriter(f))
}

func CreateORExists(FilePath string) (*os.File, error) {
	if _, err := os.Stat(FilePath); err != nil {
		if os.IsNotExist(err) {
			return os.Create(FilePath)
		}
		return nil, err
	}
	return os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY, 0600)
}
