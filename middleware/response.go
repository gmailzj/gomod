package middleware

import (
	"fmt"
	"gomod/data"
	errorcode "gomod/errorcode"
	"gomod/utils"
	"gomod/utils/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context) {
	logId := utils.Uniqid()
	c.Set(utils.LogIdParam, logId)
	logger.Info(logId, "收到请求", c.Request.RequestURI, c.Request.Method)

	c.Header("server", "golden/1.0.0")
	c.Header("logid", logId)

	t := time.Now()
	c.Set("stat_begin_ts", t.Unix())
	c.Next()

	if c.Writer.Written() {
		logger.Info(logId, "响应请求-已响应过", c.Request.RequestURI, c.Request.Method)
		return
	}

	resp, ok := c.Get("response")
	if !ok {
		logger.Error(logId, "竟然没有设置response", resp)
		c.JSON(200, gin.H{
			"code":    errorcode.InternalServerError,
			"message": errorcode.ErrorMsg[errorcode.InternalServerError],
			"data":    map[string]string{},
		})
		return
	}

	res, ok := resp.(data.Resp)
	if !ok {
		logger.Error(logId, "Resp竟然不是data.Resp类型", res)
		c.JSON(200, gin.H{
			"code":    errorcode.InternalServerError,
			"message": errorcode.ErrorMsg[errorcode.InternalServerError],
			"data":    map[string]string{},
		})
		return
	}

	var statusCode int
	if res.StatusCode == 0 {
		statusCode = 200
	} else {
		statusCode = res.StatusCode
	}

	var message string
	if res.Error == nil {
		if res.ErrorDesc != "" {
			message = res.ErrorDesc
		} else if res.ErrorParams != "" {
			message = errorcode.ErrorMsg[res.ErrorCode]
			message = strings.Replace(message, "{PARAM}", res.ErrorParams, 1)
		} else {
			message = errorcode.ErrorMsg[res.ErrorCode]
		}
	} else {
		switch e := res.Error.(type) {
		case data.InternalError:
			logger.Error(e.LogId, "服务器内部错误", "产生错误的原因", e.OriginError)
			res.ErrorCode = errorcode.InternalServerError
			message = res.Error.Error()
		case data.ValidateError:
			res.ErrorCode = errorcode.ValidateError
			message = res.Error.Error()
		case data.BusinessDataError:
			res.ErrorCode = errorcode.ValidateError
			message = res.Error.Error()
		default:
			logger.Error(logId, "忘记检验的错误类型", e, res)
		}

		if res.ErrorCode == 0 {
			res.ErrorCode = errorcode.InternalServerError
		}
		if message == "" {
			message = errorcode.ErrorMsg[errorcode.InternalServerError]
		}
	}
	if res.Data == nil {
		res.Data = map[string]string{}
	}

	if message == "" {
		message = errorcode.ErrorMsg[errorcode.Ok]
	}

	rp := gin.H{
		"code":    res.ErrorCode,
		"message": message,
		"data":    res.Data,
	}
	c.JSON(statusCode, rp)
	logger.Info(logId, "响应请求", c.Request.RequestURI, c.Request.Method, fmt.Sprintf("%+v", rp))

	//统计接口调用时间
	latency := time.Since(t)
	reuestPath := c.Request.URL.Path
	logger.Info(reuestPath, "http_client", 0, int64(latency/1e6))

	c.Abort()
	return
}
