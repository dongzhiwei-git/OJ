package services

import (
	"context"
	"fmt"
	"github.com/dongzhiwei-git/dzwlib/pkgs/terror"
	"hgoj/internal/dao"
	"hgoj/internal/models"
	"hgoj/internal/pkg"
	"log"
	"time"
)

type Problem struct {
	ProblemID    int32     `gorm:"problem_id"    json:"problem_id"`
	Title        string    `gorm:"title"         json:"title"         binding:"required"`
	Description  string    `gorm:"description"   json:"description"   binding:"required"`
	Input        string    `gorm:"input"         json:"input"         binding:"required"`
	Output       string    `gorm:"output"        json:"output"        binding:"required"`
	SampleInput  string    `gorm:"sample_input"  json:"sampleInput"   binding:"required"`
	SampleOutput string    `gorm:"sample_output" json:"sampleOutput"  binding:"required"`
	Spj          string    `gorm:"spj"           json:"spj"           binding:"required"`
	Hint         string    `gorm:"hint"          json:"hint"`
	Source       string    `gorm:"source"        json:"source"`
	InDate       time.Time `gorm:"in_date"       json:"inDate"`
	TimeLimit    int32     `gorm:"time_limit"    json:"time"          binding:"required"`
	MemoryLimit  int32     `gorm:"memory_limit"  json:"memory"        binding:"required"`
	Accepted     int32     `gorm:"accepted"      json:"accepted"`
	Submit       int32     `gorm:"submit"        json:"submit"`
	Solved       int32     `gorm:"solved"        json:"solved"`
	FileName     string    `json:"fileName"     binding:"required"`
}

// QueryAllProblem 查询所有题目
func (pr *Problem) QueryAllProblem(ctx context.Context) (*[]*models.Problem, error) {
	problem := new([]*models.Problem)
	err := dao.Orm.Find(problem).Error

	return problem, terror.Trace(fmt.Errorf("test   err: %v", err))
}

// QueryProblemByPageNum 通过页码查询题目
func (pr *Problem) QueryProblemByPageNum(num, size int) (proInfo []*models.Problem, count int64, err error) {
	start := pkg.StartNum(num, size)
	err = dao.Orm.Order("problem_id desc").Limit(size).Offset(start).Find(&proInfo).Error
	err = dao.Orm.Model(&models.Problem{}).Count(&count).Error
	return proInfo, count, err
}

// AddProblem 添加问题
func (pr *Problem) AddProblem(data models.Problem) int32 {
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

// DelProblemById 通过problem删除一行
func (pr *Problem) DelProblemById(id int32) error {
	err := dao.Orm.Debug().Where("problem_id = ?", id).Delete(&models.Problem{}).Error

	return err
}

// QueryProblemByProblemID 通过ID获取某一个题目
func (pr *Problem) QueryProblemByProblemID(id int) (pro models.Problem, err error) {
	proID := int32(id)
	err = dao.Orm.Debug().Where("problem_id = ?", proID).Find(&pro).Error
	return pro, err
}
