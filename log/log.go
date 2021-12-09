package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func sprintf(t string, format string, args ...interface{}) string {
	_, path, line, _ := runtime.Caller(3)
	_, file := filepath.Split(path)
	return fmt.Sprintf(fmt.Sprintf("%s [%s] %s:%d %s", time.Now().Format("2006-01-02 15:04:05.999"), t, file, line, format), args...)
}

func printf(t string, format string, args ...interface{}) {
	fmt.Printf(sprintf(t, format, args...))
}

func Infof(format string, args ...interface{}) {
	printf("INFO", format, args...)
}

func Info(arg interface{}) {
	printf("INFO", "%v\n", arg)
}

func Errorf(format string, args ...interface{}) {
	printf("ERROR", format, args...)
}

func Error(arg interface{}) {
	printf("ERROR", "%v\n", arg)
}

func Panicf(format string, args ...interface{}) {
	panic(sprintf("PANIC", format, args...))
}

func Panic(arg interface{}) {
	printf("PANIC", "%v\n", arg)
}
