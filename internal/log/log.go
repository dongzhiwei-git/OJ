package log

import "github.com/sirupsen/logrus"

func Init() {
	logrus.SetLevel(logrus.TraceLevel) // 在测试环境中设置低等级级别，
	//logrus.SetLevel(logrus.InfoLevel)    // 在生产环境中需要考虑性能，关注关键信息，level 设置高一点
	logrus.SetReportCaller(true) // 调用者文件名与位置
	//logrus.SetFormatter(new(logrus.JSONFormatter))    // 日志格式设置成json
	//logrus.Traceln("trace 日志")
	//logrus.Debugln("debug 日志")
	//logrus.Infoln("Info 日志")
	//logrus.Warnln("warn 日志")
	//logrus.Errorln("error msg")
	//logrus.Fatalf("fatal 日志")
	//logrus.Panicln("panic 日志")
}
