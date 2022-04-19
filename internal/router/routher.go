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
	r.GET("/problem_page", api.QueryProblemByPageNum)
	r.GET("/problem_by_proid", api.QueryProblemByProblemID)

	// 保存测试数据

	var admin = r.Group("/admin")
	{
		//上传测试数据
		admin.POST("/upload", api.FileUpload)
		admin.POST("/reg", api.CreateAdminUser)
		//添加问题
		admin.POST("/add_problem", api.AddProblem)
		admin.POST("/add_contest", api.AddContest)
	}

	// 提交代码
	r.POST("/submit", api.Submit)
	r.GET("/status", api.Status)

	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}

}
