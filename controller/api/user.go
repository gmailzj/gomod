package api

import (
	"github.com/gin-gonic/gin"
	myjwt "gomod/middleware/jwt"
	"net/http"
)

func UserInfo(c *gin.Context) {

	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": 0,
	// 	"msg":    "登录成功！",
	// 	"data":   "123",
	// })
}
