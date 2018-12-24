package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	c.String(http.StatusOK, strings.ToUpper("index"))
}
