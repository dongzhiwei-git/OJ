package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
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
