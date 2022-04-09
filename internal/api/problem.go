package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inherited/internal/services"
	"net/http"
)

func  QueryAllProblem(ctx *gin.Context)  {
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