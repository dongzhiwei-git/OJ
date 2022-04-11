package router

import (
	"fmt"
	"inherited/internal/api"
	"inherited/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine
	r = gin.Default()
	// to solve the cross domain
	r.Use(pkg.Cors())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	r.GET("/problem", api.QueryAllProblem)
	// 保存测试数据

	var admin = r.Group("/admin")
	{
		admin.POST("/upload", api.FileUpload)
		admin.POST("/reg", api.CreateAdminUser)
	}

	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}

}
