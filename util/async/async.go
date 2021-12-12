package async

import (
	"bytes"
	"fmt"
	"runtime"
)

func Async(fun func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				PrintStackTrace(err)
			}
		}()
		fun()
	}()
}

func PrintStackTrace(err interface{}) {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%v\n", err))
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buf.WriteString(fmt.Sprintf("%s:%d (0x%x)\n", file, line, pc))
	}
	fmt.Println(buf.String())
}
