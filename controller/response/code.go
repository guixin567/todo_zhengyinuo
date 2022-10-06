package response

var codeMessageMap = map[int64]string{
	CodeSuccess:                 "success",
	CodeInvalidParam:            "invalid request params",
	CodeUserNameOrPasswordError: "user name or password is wrong",
	CodeAuthInvalid:             "auth failed",
	CodeNeedLogin:               "need login",
}

const (
	CodeSuccess = 200 + iota
	CodeInvalidParam
	CodeUserNameOrPasswordError
	CodeAuthInvalid
	CodeNeedLogin
)

func getMsg(code int64) string {
	return codeMessageMap[code]
}
