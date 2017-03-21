package log

import (
	glog "log"
	"runtime"
)

func LogCodeLine() {
	if _, file, line, ok := runtime.Caller(2); ok {
		glog.Printf("%s:%d\n", file, line)
	}
}
