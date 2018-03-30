package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	chanStart := time.Now()
	runYesChan()
	chanDuration := time.Since(chanStart)

	readerStart := time.Now()
	runYesReader()
	readerDuration := time.Since(readerStart)

	log.Printf("chan took %s", chanDuration)
	log.Printf("reader took %s", readerDuration)
}

func runYesChan() {
	b := make(chan string)
	go yesChan(b, "y\n")
	c := append([]string{}, "123")
	for i := 0; i < 10000; i++ {
		fmt.Print(<-b, c)
	}
}

func runYesReader() {
	y := YesReader{[]byte("y\n")}
	io.CopyN(os.Stdout, y, 20000)
}

func yesChan(b chan string, s string) {
	for {
		b <- s
	}
}

// YesReader repeats bytes as a reader
type YesReader struct {
	template []byte
}

func (y YesReader) Read(b []byte) (int, error) {
	for i, v := range y.template {
		b[i] = v
	}
	return len(y.template), nil
}
