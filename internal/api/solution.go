package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "successful",
	})
}
