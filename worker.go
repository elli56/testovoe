package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Info struct {
	input      []string
	word       string
	amountWord chan int
	maxThreat  int
	logger     *log.Logger
}

func (i *Info) Start() {
	wgr := sync.WaitGroup{}
	wgr.Add(1)
	go i.readAndPrint(&wgr)

	tokens := make(chan struct{}, i.maxThreat) // подсчитывающий семафор
	wgw := sync.WaitGroup{}

	for _, url := range i.input {
		tokens <- struct{}{}

		wgw.Add(1)
		go i.worker(url, &wgw, tokens)
	}
	wgw.Wait()
	close(i.amountWord)

	wgr.Wait()
	// i.logger.Println("finish")
}

func (i *Info) worker(url string, wg *sync.WaitGroup, tokens chan struct{}) {
	body := i.getRequest(url)

	count := i.countWord(body)

	i.amountWord <- count

	i.logger.Printf("Count for %s: %d", url, count)

	wg.Done()

	<-tokens
}

func (i *Info) getRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		i.logger.Fatalf("get failed: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		i.logger.Fatalf("body read failed: %s", err.Error())
	}

	return string(body)
}

func (i *Info) countWord(body string) int {
	return strings.Count(body, i.word)
}

func (i *Info) readAndPrint(wg *sync.WaitGroup) {
	var count int

	for val := range i.amountWord {
		count += val
	}

	i.logger.Printf("Total: %d", count)
	wg.Done()
}
