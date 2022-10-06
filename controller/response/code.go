package response

var codeMessageMap = map[int64]string{
	CodeSuccess:                 "success",
	CodeInvalidParam:            "invalid request params",
	CodeUserNameOrPasswordError: "user name or password is wrong",
	CodeAuthInvalid:             "auth failed",
	CodeNeedLogin:               "need login",
	CodeServerError:             "server error",
}

const (
	CodeSuccess = 200 + iota
	CodeInvalidParam
	CodeUserNameOrPasswordError
	CodeAuthInvalid
	CodeNeedLogin
	CodeServerError
)

func getMsg(code int64) string {
	return codeMessageMap[code]
}
