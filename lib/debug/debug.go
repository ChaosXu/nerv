package debug

import (
	"fmt"
	"runtime"
)

func CallerLine() string {
	if _, file, line, ok := runtime.Caller(2); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}

func CodeLine() string {
	if _, file, line, ok := runtime.Caller(1); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}
