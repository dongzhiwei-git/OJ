package test

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	os.Exit(m.Run())

}

//import (
//	"inherited/internal"
//	"logs"
//)
//
//func init() {
//	// 初始化依赖模块
//	if err := internal.Init(); err != nil {
//		logs.Println("Init failed." + err.Error())
//		return
//	}
//
//}

func TestCreateSysUser(t *testing.T) {

	SysUser := new(services.SysUser)
	SysUser.CreateSysUser("rot", "2355")
	fmt.Println("er")

}

func TestQueryAllProblem(t *testing.T) {

	problem := new(services.Problem)
	info, _ := problem.QueryAllProblem()
	fmt.Println(info)

}
