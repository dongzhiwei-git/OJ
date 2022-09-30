package services

import (
	"hgoj/internal/dao"
	"hgoj/internal/models"
)

type User struct{}

// GetAllAdmin 查询所有管理员信息
func (user *User) GetAllAdmin() (*[]*models.User, error) {
	users := new([]*models.User)
	//err := dao.Orm.Find(users).Error
	err := dao.Orm.Where("role = 1").Find(users).Error
	return users, err
}
