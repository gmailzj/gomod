package errorcode

type Code int

const (
	Ok                 Code = 0
	VerError           Code = 1001
	PublicParamAbsence Code = 1002
	TimestampError     Code = 1003
	ValidateError      Code = 1004
	SignatureError     Code = 1005
	BusinessDataError  Code = 1006

	BadRequest          Code = 400
	Forbidden           Code = 403
	InternalServerError Code = 500
)

var ErrorMsg = map[Code]string{
	Ok:                 "success",
	PublicParamAbsence: "{PARAM} 不能为空",
	SignatureError:     "签名错误",
	TimestampError:     "与服务器时间误差不能大于30s",
	VerError:           "版本号错误,当前存在的版本{PARAM}",
	ValidateError:      "参数验证错误 {ERRMSG}",
	BusinessDataError:  "业务参数校验错误 {ERRMSG}",

	BadRequest:          "Bad Request",
	Forbidden:           "Forbidden",
	InternalServerError: "Server Internal Error",
}
