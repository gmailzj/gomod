package errorCode

type Code int

const (
	Ok                       Code = 0
	VerError                 Code = 1001
	PublicParamAbsence       Code = 1002
	TimestampError           Code = 1003
	AppKeyNotExists          Code = 1004
	DeveloperAppClosed       Code = 1005
	DeveloperClosed          Code = 1006
	SignatureError           Code = 1007
	DeveloperNeedAudit       Code = 1008
	DeveloperAuditFailed     Code = 1009
	ValidateError            Code = 1010
	ServerErrorNeedAdmin     Code = 1011
	LackOfAvailableCount     Code = 1012
	ChargeError              Code = 1013
	BusinessDataError        Code = 1014
	ExternalDependencyReject Code = 1015
	ExternalDependencyError  Code = 1016

	BadRequest          Code = 400
	Forbidden           Code = 403
	InternalServerError Code = 500
)

var ErrorMsg = map[Code]string{
	Ok:                       "success",
	PublicParamAbsence:       "{PARAM} 不能为空",
	SignatureError:           "签名错误",
	TimestampError:           "与服务器时间误差不能大于30s",
	AppKeyNotExists:          "appkey不存在",
	DeveloperNeedAudit:       "开发者待审核",
	DeveloperAuditFailed:     "开发者审核未通过",
	ExternalDependencyError:  "依赖其他项目错误 {ERRMSG}",
	ServerErrorNeedAdmin:     "内部错误，请联系管理员并提供LogId {LOGID}",
	VerError:                 "版本号错误,当前存在的版本{PARAM}",
	ValidateError:            "参数验证错误 {ERRMSG}",
	ExternalDependencyReject: "{PARAM}",
	LackOfAvailableCount:     "无可用接口调用次数,请购买",
	ChargeError:              "计费异常,请联系管理员",
	DeveloperAppClosed:       "app已经被关闭",
	DeveloperClosed:          "开发者已经被关闭",
	BusinessDataError:        "业务参数校验错误 {ERRMSG}",

	BadRequest:          "Bad Request",
	Forbidden:           "Forbidden",
	InternalServerError: "Server Internal Error",
}
