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
func (pr *Problem) AddProblem(data []string, inDate time.Time) int64 {
	var pro models.Problem
	//proo :=  new(models.Problem)
	pro.Title = data[0]
	pro.TimeLimit = stringToInt32(data[1])
	pro.MemoryLimit = stringToInt32(data[2])
	pro.Description = data[3]
	pro.Input = data[4]
	pro.Output = data[5]
	pro.SampleInput = data[6]
	pro.SampleOutput = data[7]
	pro.Hint = data[8]
	pro.Spj = data[9]
	pro.InDate = inDate
	pro.Defunct = "N"

	err := dao.Orm.Debug().Create(&pro).Error
	if err != nil {
		log.Println("[service.AddProblem]", err)
		return 0

	}

	err = dao.Orm.Debug().Last(&pro).Error
	if err != nil {
		log.Println("[service.AddProblem]", err)
		return 0
	}

	return pro.ProblemID
}
