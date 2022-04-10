package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inherited/internal/services"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	OJ_DATA          = "/home/judge/data"
	OJ_ZIP_TEMP_DATA = OJ_DATA + "/tempzip"
	SESSION_USER_KEY = "dongzhiwei"
)

func QueryAllProblem(ctx *gin.Context) {
	problem := new(services.Problem)
	pro_set, err := problem.QueryAllProblem()
	if err != nil {
		fmt.Printf("[api.QueryAllProblem], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": pro_set,
	})

	return

}

// FileUpload 上传测试数据zip
func FileUpload(ctx *gin.Context) {

	key := time.Now().String()
	//f, h, err := this.GetFile("file")
	zipDir := OJ_ZIP_TEMP_DATA
	file, err := os.OpenFile(zipDir, os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Fatal("[api.FileUpload]", err)
		return
	}

	err2 := os.Mkdir(zipDir, os.ModePerm)
	if err2 != nil {
		//logs.Error(err2)
	}

	//err1 := this.SaveToFile("file", OJ_ZIP_TEMP_DATA+"/"+key+h.Filename)
	//if err1 != nil {
	//	this.JsonErr("文件上传错误", 24005, "")
	//}

}
