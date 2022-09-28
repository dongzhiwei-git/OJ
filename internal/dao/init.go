package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hgoj/internal/conf"
)

var Orm *gorm.DB

func Init() (err error) {
	// 初始化db

	conn := conf.Get().Mysql
	//DB, err := gorm.Open("mysql", "root:beego@tcp(121.36.216.191:3306)/inherited?charset=utf8mb4&parseTime=True&loc=Local")
	Orm, err = gorm.Open(mysql.Open(conn.Master.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//Orm.AutoMigrate(models.Problem{})
	//Orm.AutoMigrate(models.SysUser{})
	//SysUser := new(services.SysUser)
	//SysUser.CreateSysUser("root", "234")
	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}

	return err

	// 初始化oss3

	return
}
