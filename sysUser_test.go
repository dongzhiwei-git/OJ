package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hgoj/internal"
	"hgoj/internal/models"
	"hgoj/internal/services"

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

func TestAddProblem(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	problem := new(services.Problem)
	data := models.Problem{}
	info := problem.AddProblem(data)
	fmt.Println(info)

}

func TestDelProblemById(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	pro := new(services.Problem)
	pro.DelProblemById(3)
}

func TestAddSolution(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	sol := new(services.Solution)
	_, err := sol.AddSolution(34, "!23 ", 1, 1, 34, 1, "3")
	if err != nil {
		fmt.Println()
	}

}

func TestQueryContestByConId(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	con := new(services.Contest)
	conID, err := con.QueryContestByConId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conID)
}

func TestStatus(t *testing.T) {
	solu := new(services.Solution)
	soluInfos, count, err := solu.GetStatusPage(3, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soluInfos, "\n", count)
}

func TestQueryProblemByPageNum(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	pro := new(services.Problem)
	proInfo, count, err := pro.QueryProblemByPageNum(3, 2)
	if err != nil {
		fmt.Println(err)
		logrus.Trace(err)
	}
	fmt.Println(proInfo, count)

}

func TestQueryProblemByProblemID(t *testing.T) {
	pro := new(services.Problem)
	proInfo, err := pro.QueryProblemByProblemID(27)
	if err != nil {

		logrus.Info("[数据库查询失败]")

		return
	}
	fmt.Println(proInfo)
}
