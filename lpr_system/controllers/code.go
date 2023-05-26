package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeServerBusy
	CodeLogin
	CodeEdit
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "sucess",
	CodeInvalidParam: "请求参数错误",
	CodeServerBusy:   "服务器繁忙",
	CodeLogin:        "failed",
	CodeEdit:         "该学生已提交",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
