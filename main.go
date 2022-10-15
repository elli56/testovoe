package main

import (
	"log"
)

func main() {
	input := []string{"https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org", "https://golang.org"}
	word := "Go"

	logger := log.Default()

	inf := Info{
		input:  input,
		word:   word,
		logger: logger,
	}

	inf.Worker()
}
