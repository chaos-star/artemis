package Code

import "runtime"

type Exception struct {
	code       int
	methodName string
	fileName   string
	line       int
	err        error
}

func (e *Exception) SetCode(code int, err error) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		e.fileName = file
		e.line = line
		e.methodName = runtime.FuncForPC(pc).Name()
	}
	e.code = code
	e.err = err
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) MethodName() string {
	return e.methodName
}

func (e *Exception) FileName() string {
	return e.fileName
}

func (e *Exception) Line() int {
	return e.line
}

func (e *Exception) Error() string {
	message := ""
	if e.err != nil {
		message = e.err.Error()
	}
	return message
}

func (e *Exception) IsError() bool {
	var flag bool
	if e.code != 0 {
		flag = true
	}
	return flag
}
