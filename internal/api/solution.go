package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"inherited/internal/dhttp"
	"inherited/internal/services"
	"net/http"
	"strconv"
)

// 分页得到判题页
func Status(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	if pageSize == 0 || pageNum == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.ParaInvaild,
			"msg":  "参数无效",
		})
		logrus.Info("[参数无效]")
		return
	}
	solu := new(services.Solution)
	soluInfo, count, err := solu.GetStatusPage(pageNum, pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.DatabaseRError,
		})
		logrus.Info("[数据库查询失败]")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "successful",
		"data":  &soluInfo,
		"count": count,
	})
}
