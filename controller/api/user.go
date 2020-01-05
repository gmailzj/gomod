package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   "123",
	})
}
