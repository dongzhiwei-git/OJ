package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hgoj/internal/services"
	"net/http"
)

// GetAllAdmin 查询所有管理员信息
func GetAllAdmin(ctx *gin.Context) {
	User := new(services.User)
	users, err := User.GetAllAdmin()
	if err != nil {
		fmt.Printf("[api.CheckAllUserByAdmin], err: %v", err)

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
