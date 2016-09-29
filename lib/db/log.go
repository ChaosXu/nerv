package db

import (
	"fmt"
	"runtime"
)

func logCodeLine() {
	if fn, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("%s:%d:%s\n", file, line, runtime.FuncForPC(fn).Name())
	}
}

