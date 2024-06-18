package main

import (
	"log"
	"time"
)

func main() {
	for true {
		time.Sleep(time.Second * 3)
		log.Println("2")
	}
}

//编译
//	go build -o build/main main.go
//运行
//  ./build/main > log.log 2>&1
//查看日志在打印1
//改代码打印2,再次编译
//日志依旧打印1
//找到运行进程
//	ps -ef |grep build/main
//kill进程
//  kill 29634
