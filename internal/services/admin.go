package services

import (
	"hgoj/internal/dao"
	"hgoj/internal/models"
)

type SysUser struct {
}

func (s *SysUser) CreateSysUser(name, password string) (err error) {

	var adminUser = models.SysUser{

		UserName: name,
		Password: password,
	}

	err = dao.Orm.Create(&adminUser).Error

	return
}
