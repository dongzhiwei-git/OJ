package models

import "time"

var JUDGESTATUS = map[int16]string{
	0:  "判题中",
	1:  "等待重判",
	2:  "编译中",
	3:  "运行并评判",
	4:  "正确",
	5:  "格式错误",
	6:  "答案错误",
	7:  "时间超限",
	8:  "内存超限",
	9:  "输出超限",
	10: "运行错误",
	11: "编译错误",
	12: "编译成功",
	13: "运行完成",
}

var JUDGERESCLSAA = map[int]string{
	0:  "warning",
	1:  "danger",
	2:  "warning",
	3:  "info",
	4:  "success",
	5:  "warning",
	6:  "danger",
	7:  "warning",
	8:  "warning",
	9:  "warning",
	10: "danger",
	11: "danger",
	12: "info",
	13: "success",
}

type Solution struct {
	SolutionId int32
	ProblemId  int32
	UserId     int32
	Time       int32
	Memory     int32
	InDate     time.Time
	Result     int16
	Language   uint
	Ip         string
	ContestId  int32
	Valid      int8
	Num        int8
	CodeLength int32
	Judgetime  *time.Time
	PassRate   float64
	LintError  uint
	Judger     string
}

func (Solution) TableName() string {
	return "solution"
}
