package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	// 获取 --log_dir 标志的值
	logDir := flag.Lookup("log_dir")
	if logDir != nil {
		glog.Infof("Log directory: %s\n", logDir.Value.String())
	} else {
		glog.Infof("Log directory not set")
	}

	glog.Info("This is an info message")
	glog.Warning("This is a warning message")
	glog.Error("This is an error message")
}

// go run main.go
//E1008 11:51:09.187285   93385 main.go:15] This is an error message

// go run main.go --alsologtostderr
//I1008 11:51:21.904075   93484 main.go:13] This is an info message
//W1008 11:51:21.905074   93484 main.go:14] This is a warning message
//E1008 11:51:21.905755   93484 main.go:15] This is an error message

// go run main.go --log_dir=./log

// go run main.go --alsologtostderr -v=2 -stderrthreshold=INFO --log_dir=./log
//I1008 11:51:35.499688   93585 main.go:13] This is an info message
//W1008 11:51:35.500851   93585 main.go:14] This is a warning message
//E1008 11:51:35.501366   93585 main.go:15] This is an error message

//-v=0: 仅输出最重要的信息，如错误和警告。这是默认级别，通常用于生产环境，确保只记录关键日志。
//-v=1: 输出一般的信息和警告，适合用于追踪应用程序的主要流程和状态变化。
//-v=2: 输出详细的调试信息，适合用于开发和调试阶段，帮助开发者了解应用程序的内部状态和执行流程。
//-v=3 及更高: 输出非常详细的调试信息，通常用于深入调试或诊断复杂问题。随着数字的增加，日志的详细程度也会增加。
