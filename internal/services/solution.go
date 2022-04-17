package services

import (
	"github.com/sirupsen/logrus"
	"inherited/internal/dao"
	"inherited/internal/models"
	"inherited/internal/pkg"
	"time"
)

type Solution struct {
}

// AddSolution 提交代码
func (s *Solution) AddSolution(pid int32, source string, uid int32, codeLen int, lang int, conid int32, ip string) (int64, error) {
	var Solu models.Solution
	var SoluCode models.SourceCode
	//err := DB.Begin()
	Solu.ProblemId = pid
	Solu.UserId = uid
	Solu.InDate = time.Now()
	Solu.Language = uint(lang)
	Solu.CodeLength = int32(codeLen)
	Solu.Result = 0
	logrus.Info("conid:", conid)
	if conid != -1 {
		Solu.ContestId = conid
	}
	Solu.Ip = ip

	//id, err := DB.Insert(&Solu)
	tx := dao.Orm.Begin()
	err := tx.Create(&Solu).Error
	//获取插入记录的Id
	var sid int
	err = tx.Raw("select LAST_INSERT_ID() as id").Pluck("id", &sid).Error
	SoluCode.SolutionId = int32(sid)
	SoluCode.Source = source

	err = tx.Create(&SoluCode).Error
	var scid int
	err = tx.Raw("select LAST_INSERT_ID() as id").Pluck("id", &scid).Error

	if sid == 0 || scid == 0 || err != nil {
		err = tx.Rollback().Error
		logrus.Fatal("数据库回滚失败")
		return int64(sid), err
	} else {
		err = tx.Commit().Error
		if err != nil {
			logrus.Fatal("事务提交失败")
		}

		return int64(sid), err
	}

}

// GetStatusPage 分页获取判题信息
func (s Solution) GetStatusPage(PageNum, pageSize int) (solInfo []models.Solution, tot int64, err error) {
	startNum := pkg.StartNum(PageNum, pageSize)
	err = dao.Orm.Debug().Limit(pageSize).Offset(startNum).Find(&solInfo).Error
	if err != nil {
		logrus.Info("[Err]")
	}
	err = dao.Orm.Debug().Model(&models.Solution{}).Count(&tot).Error
	if err != nil {
		logrus.Info("[Err]")
	}

	return solInfo, tot, err
}
