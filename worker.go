package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Info struct {
	input     []string
	word      string
	countWord int
	logger    *log.Logger
}

func (i *Info) Worker() {
	var count int

	for _, url := range i.input {
		bodyResp, err := i.getRequest(url)
		if err != nil {
			i.logger.Println(err.Error())
		}

		count = strings.Count(bodyResp, i.word)

		i.logger.Printf("count: %d", count)

		i.countWord += count
	}

	i.logger.Printf("full count: %d", i.countWord)
}

func (i *Info) getRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
