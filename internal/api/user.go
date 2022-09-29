package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hgoj/internal/services"
	"net/http"
)

// CheckAllUserByAdmin 管理员查询所有用户信息
func CheckAllUserByAdmin(ctx *gin.Context) {
	User := new(services.User)
	users, err := User.CheckAllUserByAdmin()
	if err != nil {
		fmt.Printf("[api.QueryAllProblem], err: %v", err)

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
