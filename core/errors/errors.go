package errors

import (
	"runtime"
)

type (
	// Frame .
	Frame struct {
		// Func contains a function name.
		Func string
		// Line contains a line number.
		Line int
		// Path contains a file path.
		Path string
	}
	// Error .
	Error interface {
		Error() string
		StackTrace() []Frame
		Unwrap() error
	}
)

// ErrorStackTrace .
// skip Es la posicion de donde comienza el tracking del error
// El valor de skip debe empezar en "2" para no tomar en cuenta esta funcion "errorStackTrace()"
// Con esto tenemos un mejor y claro tracking
func ErrorStackTrace(skip int) []Frame {
	frames := make([]Frame, 0, 10)
	for {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok || skip == 5 {
			break
		}
		fn := runtime.FuncForPC(pc)
		frame := Frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}
	return frames
}
