package models

type SysUser struct {
	Id       int    `gorm:"id" json:"id" form:"id"`
	UserName string `gorm:"user_name" json:"user_name" from:"user_name"`
	Password string `gorm:"password" json:"password" from:"password"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
