package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"inherited/internal/models"
	"inherited/internal/pkg"
	"inherited/internal/services"
	"log"
	"net/http"
	"os"
	"strconv"
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

	file, err := ctx.FormFile("file")
	key := time.Now().Format("2006-01-02 15:04:05")
	key = pkg.MD5(key)
	zipDir := OJ_ZIP_TEMP_DATA + "/" + key + file.Filename

	if err != nil {
		log.Println("[api.FileUpload]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败, name因用file"})
		return
	}

	// 保存文件
	//dir := "/Users/apple/test/" + key + file.Filename
	err = ctx.SaveUploadedFile(file, zipDir)

	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "上传成功",
		"key":      key,
		"fileName": file.Filename,
	})
}

// AddProblem 添加问题
func AddProblem(ctx *gin.Context) {
	data := models.Problem{}
	data2 := services.Problem{}
	problem := new(services.Problem)
	err := ctx.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		erro := fmt.Sprintf("%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": erro,
		})

		log.Println("[api.AddProblem1]", err)
		return
	}
	err = ctx.ShouldBindBodyWith(&data2, binding.JSON)
	if err != nil {
		erro := fmt.Sprintf("%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": erro,
		})

		log.Println("[api.AddProblem2]", err)
		return
	}

	// TODO 开启事务，失败回滚
	problemID := problem.AddProblem(data)
	fmt.Printf("problem_ID:", problemID)
	if problemID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统错误",
		})

		return
	}

	fileName := data2.FileName
	targetPath := strconv.Itoa(int(problemID))
	testZipDir := OJ_ZIP_TEMP_DATA + "/" + fileName
	targetDir := OJ_DATA + "/" + targetPath + "/"
	err = pkg.UnZip(testZipDir, targetDir)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统错误",
		})

		err = problem.DelProblemById(problemID)
		if err != nil {
			log.Println("problem delete err")
		}
		log.Println("[api.AddProblem2]", err)
		return
	}

	err = os.Remove(testZipDir)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统错误",
		})

		err = problem.DelProblemById(problemID)
		if err != nil {
			log.Println("dir delete err")
		}
		log.Println("[api.AddProblem2]", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "添加题目成功",
	})

}

// Submit 提交源代码
func Submit(ctx *gin.Context) {
	param := make(map[string]interface{})
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	if len(param["source"].(string)) == 0 {
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"msg": "代码不能为空"})
			return
		}
	}

	if param["problemID"] == 0 {
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"msg": "problemID不能为0"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
	})
}
