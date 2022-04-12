package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Problem struct {

}

//QueryAllProblem 查询所有题目
func (pr *Problem)QueryAllProblem() (*[]*models.Problem, error) {
	problem := new([]*models.Problem)
	err := dao.Orm.Find(problem).Error

	return problem,  err
}

//func AddProblem(data []string, inDate time.Time) (int64, error) {
//	var pro Problem
//	pro.Title = data[0]
//	pro.TimeLimit = stringToint32(data[1])
//	pro.MemoryLimit = stringToint32(data[2])
//	pro.Description = data[3]
//	pro.Input = data[4]
//	pro.Output = data[5]
//	pro.SampleInput = data[6]
//	pro.SampleOutput = data[7]
//	pro.Hint = data[8]
//	pro.Spj = data[9]
//	pro.InDate = inDate
//	pro.Defunct = "N"
//
//	pid, err := DB.Insert(&pro)
//	if err != nil {
//		return pid, err
//	}
//	return pid, nil
//}
