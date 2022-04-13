package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
	"strconv"
	"time"
)

type Problem struct {
}

//QueryAllProblem 查询所有题目
func (pr *Problem) QueryAllProblem() (*[]*models.Problem, error) {
	problem := new([]*models.Problem)
	err := dao.Orm.Find(problem).Error

	return problem, err
}

func stringToInt32(str string) int32 {
	d, _ := strconv.Atoi(str)
	return int32(d)
}

// AddProblem 添加问题
func (pr *Problem) AddProblem(data models.Problem) int64 {
	inDate := time.Now()
	data.InDate = inDate
	data.Defunct = "N"

	err := dao.Orm.Debug().Create(&data).Error
	if err != nil {
		log.Println("[service.AddProblem]", err)
		return 0

	}

	err = dao.Orm.Debug().Last(&data).Error
	if err != nil {
		log.Println("[service.AddProblem]", err)
		return 0
	}

	return data.ProblemID
}
