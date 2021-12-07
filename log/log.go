package log

import "fmt"

func sprintf(t string, format string, args ...interface{}) string {
	return fmt.Sprintf(fmt.Sprintf("[%s] %s", t, format), args...)
}

func printf(t string, format string, args ...interface{}) {
	fmt.Printf(sprintf(t, format, args...))
}

func Infof(format string, args ...interface{}) {
	printf("INFO", format, args...)
}

func Errorf(format string, args ...interface{}) {
	printf("ERROR", format, args...)
}

func Error(msg interface{}) {
	printf("ERROR", "%v", msg)
}

func Panicf(format string, args ...interface{}) {
	panic(sprintf("PANIC", format, args...))
}

func Panic(msg interface{}) {
	printf("PANIC", "%v", msg)
}
