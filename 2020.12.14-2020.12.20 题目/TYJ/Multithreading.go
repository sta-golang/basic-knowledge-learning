package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var ans int

var workerCount = 0
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundmatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("C:\\")
	waitForWorkers()
	fmt.Println(ans,"ans")
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path)
		case <-workerDone:
			workerCount--
			if workerCount == 0{
				return
			}
		case <-foundmatch:
			ans++
		}
	}
}


func search(path string)  {
	files,err := ioutil.ReadDir(path)
	if err == nil {
		for _,file := range files{
			name := file.Name()
			if strings.Contains(name,"a") {
				foundmatch <- true
			}
			if file.IsDir() {
				searchRequest <- path + name + "\\"
			}
		}
	}
	workerDone <- true
}
