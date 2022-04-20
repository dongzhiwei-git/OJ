package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"inherited/internal/models"
	"net/http"
)

func AddContest(ctx *gin.Context) {
	data := models.Contest{}
	err := ctx.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		erro := fmt.Sprintf("%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": erro,
		})
		logrus.Info("[", err, "]")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "successful",
		"data": data,
	})
}
