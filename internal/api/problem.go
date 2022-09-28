package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"hgoj/internal/dhttp"
	"hgoj/internal/models"
	"hgoj/internal/pkg"
	"hgoj/internal/services"

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

// QueryProblemByPageNum 通过页码查询题目
func QueryProblemByPageNum(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	if pageSize == 0 || pageNum == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.ParaInvaild,
			"msg":  "参数无效",
		})
		logrus.Info("[参数无效]")
		return
	}
	pro := new(services.Problem)
	proInfo, count, err := pro.QueryProblemByPageNum(pageNum, pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.DatabaseRError,
		})
		logrus.Info("[数据库查询失败]")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "successful",
		"data":  proInfo,
		"count": count,
	})

}

// QueryProblemByProblemID 通过ID获取某一个题目
func QueryProblemByProblemID(ctx *gin.Context) {
	problemID, _ := strconv.Atoi(ctx.Query("problemID"))
	if problemID <= 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.ParaInvaild,
			"msg":  "参数无效",
		})
		logrus.Info("[参数无效]")
		return
	}
	pro := new(services.Problem)
	proInfo, err := pro.QueryProblemByProblemID(problemID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.DatabaseRError,
		})
		logrus.Info("[数据库查询失败]")

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "successful",
		"data": proInfo,
	})

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
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		fmt.Printf("erer")
		ctx.JSON(http.StatusOK, gin.H{"err": err})
		return
	}
	source, _ := param["source"].(string)
	if len(source) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.CheckFailure,
			"msg":  "代码不能为空",
		})
		return

	}

	problemID, _ := param["problemID"].(float64)
	proID := int32(problemID)
	if proID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.CheckFailure,
			"msg":  "problemID不能为0",
		})
		logrus.Trace("problemID不能为空")
		return

	}
	userID, _ := param["userID"].(float64)
	useID := int32(userID)
	if useID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.CheckFailure,
			"msg":  "userID不能为0",
		})
		return
	}

	contestID, _ := param["contestID"].(float64)
	conID := int32(contestID)

	// 判断是不是比赛， 0不是比赛
	if conID != 0 {
		contest := new(services.Contest)
		conInfo, err := contest.QueryContestByConId(conID)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": dhttp.DatabaseRError,
			})
			logrus.Trace(err)
			return
		}
		t := time.Now().Sub(conInfo.EndTime).Seconds()
		if t > 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": dhttp.OK,
				"msg":  "比赛已经结束",
			})

			return
		}
	}

	lan, _ := param["language"].(float64)
	lang := int(lan)
	sol := new(services.Solution)
	codeLen := len(source)
	reqIP := ctx.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	_, err = sol.AddSolution(proID, source, useID, codeLen, lang, conID, reqIP)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": dhttp.DatabaseCError,
			"msg":  "数据库错误",
		})
		logrus.Trace(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
	})
}
