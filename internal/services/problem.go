package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Problem struct {

}

// QueryAllProblem 查询所有题目
func (pr *Problem)QueryAllProblem() (*[]*models.Problem, error) {
	problem := new([]*models.Problem)
	err := dao.Orm.Find(problem).Error

	return problem,  err
}
