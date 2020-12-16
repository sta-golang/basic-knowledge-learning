package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

//start := time.Now()
// time.Since(start)
//ioutil.ReadDir(path)
//file.isDir()
var cnt int
var name string = "a"

func main(){
	start := time.Now()
	time.Since(start)
	// TODO: search //
	search("E:\\")
	fmt.Printf("%d\n",cnt)
	fmt.Println(time.Since(start))
}

func search(path string){
	files,_ := ioutil.ReadDir(path)
	for _,file := range files{
		if strings.Contains(file.Name(),name){
			cnt++
		}
		if file.IsDir(){
			search(path + file.Name() + "\\")
		}
	}
}






















