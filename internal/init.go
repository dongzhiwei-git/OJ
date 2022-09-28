package internal

import (
	"hgoj/internal/conf"
	"hgoj/internal/dao"
	log2 "hgoj/internal/log"
	"log"

	"github.com/pkg/errors"
)

func Init() error {

	// 初始化配置
	if err := conf.Init(); err != nil {
		return errors.WithStack(err)
	}

	// 初始化dao
	if err := dao.Init(); err != nil {
		log.Println("初始化dao失败")
		return errors.WithStack(err)
	}

	// 初始化logrus
	log2.Init()

	return nil
}
