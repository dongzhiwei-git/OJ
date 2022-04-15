package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Contest struct {
}

// QueryContestByConId 通过contest_id查找一条记录
func (con *Contest) QueryContestByConId(contestID int32) (*models.Contest, error) {
	contestInfo := new(models.Contest)
	err := dao.Orm.Debug().First(contestInfo, contestID).Error

	return contestInfo, err

}
