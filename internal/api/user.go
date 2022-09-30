package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hgoj/internal/services"
	"net/http"
)

// GetAllUserByAdmin 管理员查询所有用户信息
func GetAllUserByAdmin(ctx *gin.Context) {
	User := new(services.User)
	users, err := User.GetAllUserByAdmin()
	if err != nil {
		fmt.Printf("[api.CheckAllUserByAdmin], err: %v", err)

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
