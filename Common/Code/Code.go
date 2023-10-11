package Code

const (
	Success       = 0
	InternalError = 100001
	UnknownError  = 100002
)

var CodeMap = map[int]string{
	Success:       "Success",
	InternalError: "Service internal error",
	UnknownError:  "Unknown error",
}

func GetMessage(code int) string {
	var message string
	if content, ok := CodeMap[code]; ok {
		message = content
	} else {
		message = CodeMap[UnknownError]
	}
	return message
}
