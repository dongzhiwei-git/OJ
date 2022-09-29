package services

import (
	"hgoj/internal/dao"
	"hgoj/internal/models"
)

type User struct{}

// CheckAllUserByAdmin 管理员查询所有用户信息
func (user *User) CheckAllUserByAdmin() (*[]*models.User, error) {
	users := new([]*models.User)
	err := dao.Orm.Find(users).Error
	return users, err
}
