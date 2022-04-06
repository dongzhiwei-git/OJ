package dao

import (
	"fmt"
	"inherited/internal/conf"
	"inherited/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func Init() (err error) {
	// 初始化db

	conn := conf.Get().Mysql
	//DB, err := gorm.Open("mysql", "root:beego@tcp(121.36.216.191:3306)/inherited?charset=utf8mb4&parseTime=True&loc=Local")
	Orm, err = gorm.Open("mysql", conn.Master.Dsn)
	// 开启打印SQL语句
	Orm.AutoMigrate(models.SysUser{})
	//SysUser := new(services.SysUser)
	//SysUser.CreateSysUser("root", "234")

	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}

	return Orm.DB().Ping()

	// 初始化oss3

	return
}
