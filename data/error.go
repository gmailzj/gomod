package data

import (
	"encoding/json"
	"gomod/errorcode"
	"gopkg.in/go-playground/validator.v8"
	"strings"
)

//服务器内部错误
func NewInternalServelError(logId string, error error) InternalError {
	return InternalError{
		LogId:       logId,
		OriginError: error,
	}
}

type InternalError struct {
	LogId       string
	OriginError error
}

func (this InternalError) Error() string {
	return this.OriginError.Error()
}

//参数检验错误
func NewValidateError(logId string, error error) ValidateError {
	return ValidateError{
		LogId:       logId,
		OriginError: error,
	}
}

type ValidateError struct {
	LogId       string
	OriginError error
}

func (this ValidateError) Error() string {
	errs, ok := this.OriginError.(validator.ValidationErrors)
	if ok {
		var errDesc = ""
		var errMessage = ""
		for _, err := range errs {
			switch err.Tag {
			case "required":
				errMessage = "不能为空"
			case "min":
				errMessage = "不能小于" + err.Param + "个长度"
			case "max":
				errMessage = "不能大于" + err.Param + "个长度"
			case "len":
				errMessage = "长度只能为" + err.Param + "个长度"
			case "eq":
				errMessage = "只能为" + err.Param + "个长度"
			case "neq":
				errMessage = "不能为" + err.Param + "个长度"
			case "oneof":
				errMessage = "只能为" + err.Param + "其中的一个"
			case "gt":
				errMessage = "只能大于" + err.Param + "个长度"
			case "gte":
				errMessage = "只能大于等于" + err.Param + "个长度"
			case "lt":
				errMessage = "只能小于" + err.Param + "个长度"
			case "lte":
				errMessage = "只能小于等于" + err.Param + "个长度"
			case "alpha":
				errMessage = "只能为字母"
			case "email":
				errMessage = "邮箱格式错误"
			case "url":
				errMessage = "网址格式错误"
			case "numeric":
				errMessage = "只能为数字"
			case "len|len":
				errMessage = "长度错误"
			default:
				errMessage = "格式错误"
			}
			errDesc += (err.Name) + errMessage + ";"
		}
		m := errDesc[0 : len(errDesc)-1]
		return strings.Replace(errorcode.ErrorMsg[errorcode.ValidateError], "{ERRMSG}", m, 1)
	}
	if je, ok := this.OriginError.(*json.UnmarshalTypeError); ok {
		m := je.Field + "类型为" + je.Type.String() + "而不是" + je.Value
		return strings.Replace(errorcode.ErrorMsg[errorcode.ValidateError], "{ERRMSG}", m, 1)
	}
	m := "其他验证错误" + this.OriginError.Error()
	return strings.Replace(errorcode.ErrorMsg[errorcode.ValidateError], "{ERRMSG}", m, 1)
}

//业务参数校验错误
func NewBusinessDataError(logId string, err string) BusinessDataError {
	return BusinessDataError{
		LogId:     logId,
		ErrorDesc: err,
	}
}

type BusinessDataError struct {
	LogId     string
	ErrorDesc string
}

func (this BusinessDataError) Error() string {
	return strings.Replace(errorcode.ErrorMsg[errorcode.BusinessDataError], "{ERRMSG}", this.ErrorDesc, 1)
}
