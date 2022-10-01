package internal

import (
	log2 "github.com/dongzhiwei-git/dzwlib/pkgs/log"
	"github.com/sirupsen/logrus"
	"hgoj/internal/conf"
	"hgoj/internal/dao"
	"log"

	"github.com/pkg/errors"
)

func Init() error {

	// 初始化配置
	conf, err := conf.Init()
	if err != nil {
		return errors.WithStack(err)
	}

	// 初始化日志
	level := logrus.InfoLevel
	if conf.Debug {
		level = logrus.DebugLevel
	}
	log2.Setup("./log", "oj", level, log2.DefaultRotateOption)
	// 初始化dao
	if err := dao.Init(); err != nil {
		log.Println("初始化dao失败")
		return errors.WithStack(err)
	}

	// 初始化logrus
	return nil
}
