package logger

import (
	"testing"
)

func TestDebugf(t *testing.T) {
	Debug().Printf("===debug: %d", 2333)
	Info().Printf("===info: %d", 666)
	Warning().Printf("===warning: %d", 998)
	Error().Printf("===error: %d", 8848)
}

func TestLogger(t *testing.T) {
	Debugf("===debugf: %s", "测试内容2333")
}
