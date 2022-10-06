package common

var codeMessageMap = map[int64]string{
	CodeSuccess:                 "success",
	CodeInvalidParam:            "invalid request params",
	CodeUserNameOrPasswordError: "user name or password is wrong",
}

const (
	CodeSuccess = 200 + iota
	CodeInvalidParam
	CodeUserNameOrPasswordError
)

func getMsg(code int64) string {
	return codeMessageMap[code]
}
