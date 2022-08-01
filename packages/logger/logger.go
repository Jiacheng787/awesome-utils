package logger

import (
	"log"
	"os"
)

const (
	flag       = log.LstdFlags | log.Lshortfile
	preDebug   = "[DEBUG] "
	preInfo    = "[INFO] "
	preWarning = "[WARNING] "
	preError   = "[ERROR] "
)

// log 模块提供了默认预设，例如 log.Print() 和 log.Fatal()
// 如果需要修改默认配置，可以通过 log.New() 定制

// 在 golang 中 const 只能用于基本类型，对象、数组只能用 var 声明
var logger = log.New(os.Stdout, preDebug, flag)

var (
	Debugf   = log.New(os.Stdout, preDebug, flag).Printf
	Infof    = log.New(os.Stdout, preInfo, flag).Printf
	Warningf = log.New(os.Stdout, preWarning, flag).Printf
	Errorf   = log.New(os.Stdout, preError, flag).Printf
)

func Debug() *log.Logger {
	logger.SetPrefix(preDebug)
	return logger
}

func Info() *log.Logger {
	logger.SetPrefix(preInfo)
	return logger
}

func Warning() *log.Logger {
	logger.SetPrefix(preWarning)
	return logger
}

func Error() *log.Logger {
	logger.SetPrefix(preError)
	return logger
}

func compile() {
	flag := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(os.Stdout, "[INFO] ", flag)

	logger.Printf("===测试内容：%v", 2333)
}
