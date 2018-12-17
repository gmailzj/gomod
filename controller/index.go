package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, strings.ToUpper("index"))
}