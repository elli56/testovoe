package main

import (
	"log"
)

func main() {
	input := []string{
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://habr.com/ru/post/490336/",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://habr.com/ru/post/490336/",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://habr.com/ru/post/490336/",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://habr.com/ru/post/490336/",
	}

	word := "Go"

	logger := log.Default()

	inf := Info{
		input:      input,
		word:       word,
		amountWord: make(chan int),
		maxThreat:  5,
		logger:     logger,
	}

	inf.Start()
}
