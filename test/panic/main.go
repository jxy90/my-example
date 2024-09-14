package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

func init() {
	// 创建日志文件
	logFile, err := os.OpenFile("panic1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}

	// 设置日志输出到文件
	log.SetOutput(logFile)
}

func main() {
	// 捕获主函数中的 panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic occurred: %v\n", r)
			log.Printf("Stack trace: %s\n", debug.Stack())
		}
	}()

	// 模拟一个会导致 panic 的操作
	causePanic()
}

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			logPanic(r)
			//log.Printf("Panic in causePanic: %v\n", r)
			//log.Printf("Stack trace in causePanic: %s\n", debug.Stack())
		}
	}()
	// 这里故意引发 panic
	var a []int
	fmt.Println(a[1])
}

func logPanic(r interface{}) {
	// 获取堆栈信息
	stackBuf := make([]byte, 1024)
	stackBuf = stackBuf[:runtime.Stack(stackBuf, false)]

	// 格式化为单行输出
	stackStr := strings.ReplaceAll(string(stackBuf), "\n", " | ")
	logMessage := fmt.Sprintf("Panic: %v | Stack: %s", r, stackStr)
	log.Println(logMessage)
}
