package api

import (
	"fmt"
	"log"
	"net/http"

	"inherited/internal/models"
	"inherited/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateAdminUser(ctx *gin.Context) {
	//Parameter parsing
	adminUser := models.SysUser{}
	err := ctx.BindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
	}

	userName := adminUser.UserName
	password := adminUser.Password
	if userName == "" || password == "" {
		log.Printf("userName or password is null")

		return
	}

	sysUser := new(services.SysUser)
	err = sysUser.CreateSysUser(adminUser.UserName, adminUser.Password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date": "success",
	})

	return
}
