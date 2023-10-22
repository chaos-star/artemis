package Code

const (
	Success       = 0
	InternalError = 100001
	UnknownError  = 100002
	DbCreateError = 100003
	DbModifyError = 100004
	DbSelectError = 100005

	BusinessException = 200001
	UserAuthFail      = 200002
	ParamsException   = 200003
)

var CodeMap = map[string]map[int]string{
	"zh": Map_Zh,
	"en": Map_En,
	"es": Map_Es,
}

func GetMessage(code int, lang string) string {
	var message string
	if code >= 0 {
		var (
			tranCodeMap map[int]string
			ok          bool
		)
		if tranCodeMap, ok = CodeMap[lang]; !ok {
			tranCodeMap = CodeMap["zh"]
		}

		if content, has := tranCodeMap[code]; has {
			message = content
		} else {
			message = tranCodeMap[UnknownError]
		}

	}
	return message
}
