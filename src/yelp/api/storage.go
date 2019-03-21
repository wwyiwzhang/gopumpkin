package yelper

import (
	csv "encoding/csv"
	"fmt"
	"os"
)

type WriterInterface interface {
	Write(v interface{}) error
}

type csvWriter struct {
	Writer *csv.Writer
}

func newCSVWriter(f *os.File) *csvWriter {
	return &csvWriter{
		Writer: csv.NewWriter(f),
	}
}

func Writer(v interface{}) WriterInterface {
	switch f := v.(type) {
	case *os.File:
		return newCSVWriter(f)
	default:
		return nil
	}
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

func (cw *csvWriter) Write(v interface{}) error {
	var err error
	switch v.(type) {
	case []string:
		err = cw.Writer.Write(v.([]string))
	case [][]string:
		err = cw.Writer.WriteAll(v.([][]string))
	default:
		return fmt.Errorf("%T not accepted here", v)
	}
	if err != nil {
		return fmt.Errorf("failed to write to CSV: %v", err)
	}
	return nil
}
