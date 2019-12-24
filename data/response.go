package data

import "gomod/errorcode"

type Resp struct {
	StatusCode  int
	ErrorCode   errorcode.Code
	ErrorParams string
	ErrorDesc   string
	Error       error
	Data        interface{}
}
