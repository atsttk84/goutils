package trace

import (
	"fmt"
	"runtime"
)

func Trace(message string, parameters ...interface{}) (str string) {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	str = fmt.Sprintf("message=[%s] call=[%s] parameters=%+v", message, f.Name(), parameters)
	return
}

func TraceLine(message string, parameters ...interface{}) (str string) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	str = fmt.Sprintf("message=[%s] call=[%s] parameters=%+v line=[%s:%d]", message, f.Name(), parameters, file, line)
	return
}
