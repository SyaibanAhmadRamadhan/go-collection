package collection

import (
	"fmt"
	"runtime"
)

func Err(err error) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err)
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
