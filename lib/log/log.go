package log

import (
	"fmt"
	"runtime"
)

func LogCodeLine() {
	if fn, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("%s:%d:%s\n", file, line, runtime.FuncForPC(fn).Name())
	}
}

