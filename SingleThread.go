package main

import (
	"io/ioutil"
	"strings"
)

//计数器
var Ans int

func Search(path string)  {
	//读取目录并返回排好序的文件以及子目录名
	files, err := ioutil.ReadDir(path)
	//没有错误
	if err == nil {
		//遍历子目录
		for _,file := range files {
			name := file.Name()
			//查找文件名中是否包含a
			if strings.Contains(name,"a"){
				Ans++
				//println(name)
			}
			//如果file是一个文件夹 则递归调用
			if file.IsDir(){
				Search(path + name + "\\" )
			}
		}
	}
}


